package grpc

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/dymensionxyz/dymint/gerr"

	"github.com/libp2p/go-libp2p/core/crypto"
	tmp2p "github.com/tendermint/tendermint/p2p"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	rollapptypes "github.com/dymensionxyz/dymension/v3/x/rollapp/types"
	"github.com/dymensionxyz/dymint/da"
	"github.com/dymensionxyz/dymint/settlement"
	"github.com/dymensionxyz/dymint/types"

	"github.com/tendermint/tendermint/libs/pubsub"

	slmock "github.com/dymensionxyz/dymint/settlement/grpc/mockserv/proto"
)

// LayerClient is an extension of the base settlement layer client
// for usage in tests and local development.
type LayerClient struct {
	*settlement.BaseLayerClient
}

var _ settlement.LayerI = (*LayerClient)(nil)

// Init initializes the mock layer client.
func (m *LayerClient) Init(config settlement.Config, pubsub *pubsub.Server, logger types.Logger, options ...settlement.Option) error {
	HubClientMock, err := newHubClient(config, pubsub, logger)
	if err != nil {
		return err
	}
	baseOptions := []settlement.Option{
		settlement.WithHubClient(HubClientMock),
	}
	if options == nil {
		options = baseOptions
	} else {
		options = append(baseOptions, options...)
	}
	m.BaseLayerClient = &settlement.BaseLayerClient{}
	err = m.BaseLayerClient.Init(config, pubsub, logger, options...)
	if err != nil {
		return err
	}
	return nil
}

var _ settlement.HubClient = (*HubGrpcClient)(nil)

type HubGrpcClient struct {
	ctx            context.Context
	ProposerPubKey string
	slStateIndex   uint64
	logger         types.Logger
	pubsub         *pubsub.Server
	latestHeight   atomic.Uint64
	conn           *grpc.ClientConn
	sl             slmock.MockSLClient
	stopchan       chan struct{}
	refreshTime    int
}

func newHubClient(config settlement.Config, pubsub *pubsub.Server, logger types.Logger) (*HubGrpcClient, error) {
	ctx := context.Background()

	latestHeight := uint64(0)
	slStateIndex := uint64(0)
	proposer, err := initConfig(config)
	if err != nil {
		return nil, err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	logger.Debug("GRPC Dial ", "ip", config.SLGrpc.Host)

	conn, err := grpc.Dial(config.SLGrpc.Host+":"+strconv.Itoa(config.SLGrpc.Port), opts...)
	if err != nil {
		logger.Error("grpc sl connecting")
		return nil, err
	}

	client := slmock.NewMockSLClient(conn)
	stopchan := make(chan struct{})

	index, err := client.GetIndex(ctx, &slmock.SLGetIndexRequest{})
	if err == nil {
		slStateIndex = index.GetIndex()
		var settlementBatch rollapptypes.MsgUpdateState
		batchReply, err := client.GetBatch(ctx, &slmock.SLGetBatchRequest{Index: slStateIndex})
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(batchReply.GetBatch(), &settlementBatch)
		if err != nil {
			return nil, errors.New("error unmarshalling batch")
		}
		latestHeight = settlementBatch.StartHeight + settlementBatch.NumBlocks - 1
	}
	logger.Debug("Starting grpc SL ", "index", slStateIndex)

	ret := &HubGrpcClient{
		ctx:            ctx,
		ProposerPubKey: proposer,
		logger:         logger,
		pubsub:         pubsub,
		slStateIndex:   slStateIndex,
		conn:           conn,
		sl:             client,
		stopchan:       stopchan,
		refreshTime:    config.SLGrpc.RefreshTime,
	}
	ret.latestHeight.Store(latestHeight)
	return ret, nil
}

func initConfig(conf settlement.Config) (proposer string, err error) {
	if conf.KeyringHomeDir == "" {
		if conf.ProposerPubKey != "" {
			proposer = conf.ProposerPubKey
		} else {
			_, proposerPubKey, err := crypto.GenerateEd25519Key(rand.Reader)
			if err != nil {
				return "", err
			}
			pubKeybytes, err := proposerPubKey.Raw()
			if err != nil {
				return "", err
			}

			proposer = hex.EncodeToString(pubKeybytes)
		}
	} else {
		proposerKeyPath := filepath.Join(conf.KeyringHomeDir, "config/priv_validator_key.json")
		key, err := tmp2p.LoadOrGenNodeKey(proposerKeyPath)
		if err != nil {
			return "", err
		}
		proposer = hex.EncodeToString(key.PubKey().Bytes())
	}

	return
}

// Start starts the mock client
func (c *HubGrpcClient) Start() error {
	c.logger.Info("Starting grpc mock settlement")

	go func() {
		tick := time.NewTicker(time.Duration(c.refreshTime) * time.Millisecond)
		defer tick.Stop()
		for {
			select {
			case <-c.stopchan:
				// stop
				return
			case <-tick.C:
				index, err := c.sl.GetIndex(c.ctx, &slmock.SLGetIndexRequest{})
				if err == nil {
					if c.slStateIndex < index.GetIndex() {
						c.logger.Info("Simulating new batch event")

						time.Sleep(10 * time.Millisecond)
						b, err := c.retrieveBatchAtStateIndex(index.GetIndex())
						if err != nil {
							panic(err)
						}
						err = c.pubsub.PublishWithEvents(context.Background(), &settlement.EventDataNewBatchAccepted{EndHeight: b.EndHeight}, settlement.EventNewBatchAcceptedList)
						if err != nil {
							panic(err)
						}
						c.slStateIndex = index.GetIndex()
					}
				}
			}
		}
	}()
	return nil
}

// Stop stops the mock client
func (c *HubGrpcClient) Stop() error {
	c.logger.Info("Stopping grpc mock settlement")
	close(c.stopchan)
	return nil
}

// PostBatch saves the batch to the kv store
func (c *HubGrpcClient) PostBatch(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch) error {
	settlementBatch := c.convertBatchtoSettlementBatch(batch, daResult)
	c.saveBatch(settlementBatch)

	time.Sleep(10 * time.Millisecond) // mimic a delay in batch acceptance
	err := c.pubsub.PublishWithEvents(context.Background(), &settlement.EventDataNewBatchAccepted{EndHeight: settlementBatch.EndHeight}, settlement.EventNewBatchAcceptedList)
	if err != nil {
		return err
	}
	return nil
}

// GetLatestBatch returns the latest batch from the kv store
func (c *HubGrpcClient) GetLatestBatch(rollappID string) (*settlement.ResultRetrieveBatch, error) {
	c.logger.Info("GetLatestBatch grpc", "index", c.slStateIndex)
	batchResult, err := c.GetBatchAtIndex(rollappID, atomic.LoadUint64(&c.slStateIndex))
	if err != nil {
		return nil, err
	}
	return batchResult, nil
}

// GetBatchAtIndex returns the batch at the given index
func (c *HubGrpcClient) GetBatchAtIndex(rollappID string, index uint64) (*settlement.ResultRetrieveBatch, error) {
	batchResult, err := c.retrieveBatchAtStateIndex(index)
	if err != nil {
		return &settlement.ResultRetrieveBatch{
			ResultBase: settlement.ResultBase{Code: settlement.StatusError, Message: err.Error()},
		}, err
	}
	return batchResult, nil
}

func (c *HubGrpcClient) GetHeightState(index uint64) (*settlement.ResultGetHeightState, error) {
	panic("hub grpc client get height state is not implemented: implement me") // TODO: impl
}

// GetSequencers returns a list of sequencers. Currently only returns a single sequencer
func (c *HubGrpcClient) GetSequencers(rollappID string) ([]*types.Sequencer, error) {
	pubKeyBytes, err := hex.DecodeString(c.ProposerPubKey)
	if err != nil {
		return nil, err
	}
	var pubKey cryptotypes.PubKey = &ed25519.PubKey{Key: pubKeyBytes}
	return []*types.Sequencer{
		{
			PublicKey: pubKey,
			Status:    types.Proposer,
		},
	}, nil
}

func (c *HubGrpcClient) saveBatch(batch *settlement.Batch) {
	c.logger.Debug("Saving batch to grpc settlement layer", "start height",
		batch.StartHeight, "end height", batch.EndHeight)
	b, err := json.Marshal(batch)
	if err != nil {
		panic(err)
	}
	// Save the batch to the next state index
	c.logger.Debug("Saving batch to grpc settlement layer", "index", c.slStateIndex+1)
	setBatchReply, err := c.sl.SetBatch(c.ctx, &slmock.SLSetBatchRequest{Index: c.slStateIndex + 1, Batch: b})
	if err != nil {
		panic(err)
	}
	if setBatchReply.GetResult() != c.slStateIndex+1 {
		panic(err)
	}

	c.slStateIndex = setBatchReply.GetResult()

	setIndexReply, err := c.sl.SetIndex(c.ctx, &slmock.SLSetIndexRequest{Index: c.slStateIndex})
	if err != nil || setIndexReply.GetIndex() != c.slStateIndex {
		panic(err)
	}
	c.logger.Debug("Setting grpc SL Index to ", "index", setIndexReply.GetIndex())
	// Save latest height in memory and in store
	c.latestHeight.Store(batch.EndHeight)
}

func (c *HubGrpcClient) convertBatchtoSettlementBatch(batch *types.Batch, daResult *da.ResultSubmitBatch) *settlement.Batch {
	settlementBatch := &settlement.Batch{
		StartHeight: batch.StartHeight,
		EndHeight:   batch.EndHeight,
		MetaData: &settlement.BatchMetaData{
			DA: &da.DASubmitMetaData{
				Height: daResult.SubmitMetaData.Height,
				Client: daResult.SubmitMetaData.Client,
			},
		},
	}
	for _, block := range batch.Blocks {
		settlementBatch.AppHashes = append(settlementBatch.AppHashes, block.Header.AppHash)
	}
	return settlementBatch
}

func (c *HubGrpcClient) retrieveBatchAtStateIndex(slStateIndex uint64) (*settlement.ResultRetrieveBatch, error) {
	c.logger.Debug("Retrieving batch from grpc settlement layer", "SL state index", slStateIndex)

	getBatchReply, err := c.sl.GetBatch(c.ctx, &slmock.SLGetBatchRequest{Index: slStateIndex})
	if err != nil {
		return nil, gerr.ErrNotFound
	}
	b := getBatchReply.GetBatch()
	if b == nil {
		return nil, gerr.ErrNotFound
	}
	var settlementBatch settlement.Batch
	err = json.Unmarshal(b, &settlementBatch)
	if err != nil {
		return nil, errors.New("error unmarshalling batch")
	}
	batchResult := settlement.ResultRetrieveBatch{
		ResultBase: settlement.ResultBase{Code: settlement.StatusSuccess, StateIndex: slStateIndex},
		Batch:      &settlementBatch,
	}
	return &batchResult, nil
}
