// Code generated by mockery v2.43.0. DO NOT EDIT.

package settlement

import (
	da "github.com/dymensionxyz/dymint/da"
	mock "github.com/stretchr/testify/mock"

	pubsub "github.com/tendermint/tendermint/libs/pubsub"

	settlement "github.com/dymensionxyz/dymint/settlement"

	types "github.com/dymensionxyz/dymint/types"
)

// MockLayerI is an autogenerated mock type for the LayerI type
type MockLayerI struct {
	mock.Mock
}

type MockLayerI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLayerI) EXPECT() *MockLayerI_Expecter {
	return &MockLayerI_Expecter{mock: &_m.Mock}
}

// GetHeightState provides a mock function with given fields: _a0
func (_m *MockLayerI) GetHeightState(_a0 uint64) (*settlement.ResultGetHeightState, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetHeightState")
	}

	var r0 *settlement.ResultGetHeightState
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*settlement.ResultGetHeightState, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint64) *settlement.ResultGetHeightState); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultGetHeightState)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLayerI_GetHeightState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeightState'
type MockLayerI_GetHeightState_Call struct {
	*mock.Call
}

// GetHeightState is a helper method to define mock.On call
//   - _a0 uint64
func (_e *MockLayerI_Expecter) GetHeightState(_a0 interface{}) *MockLayerI_GetHeightState_Call {
	return &MockLayerI_GetHeightState_Call{Call: _e.mock.On("GetHeightState", _a0)}
}

func (_c *MockLayerI_GetHeightState_Call) Run(run func(_a0 uint64)) *MockLayerI_GetHeightState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockLayerI_GetHeightState_Call) Return(_a0 *settlement.ResultGetHeightState, _a1 error) *MockLayerI_GetHeightState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLayerI_GetHeightState_Call) RunAndReturn(run func(uint64) (*settlement.ResultGetHeightState, error)) *MockLayerI_GetHeightState_Call {
	_c.Call.Return(run)
	return _c
}

// GetProposer provides a mock function with given fields:
func (_m *MockLayerI) GetProposer() *types.Sequencer {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProposer")
	}

	var r0 *types.Sequencer
	if rf, ok := ret.Get(0).(func() *types.Sequencer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Sequencer)
		}
	}

	return r0
}

// MockLayerI_GetProposer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProposer'
type MockLayerI_GetProposer_Call struct {
	*mock.Call
}

// GetProposer is a helper method to define mock.On call
func (_e *MockLayerI_Expecter) GetProposer() *MockLayerI_GetProposer_Call {
	return &MockLayerI_GetProposer_Call{Call: _e.mock.On("GetProposer")}
}

func (_c *MockLayerI_GetProposer_Call) Run(run func()) *MockLayerI_GetProposer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLayerI_GetProposer_Call) Return(_a0 *types.Sequencer) *MockLayerI_GetProposer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_GetProposer_Call) RunAndReturn(run func() *types.Sequencer) *MockLayerI_GetProposer_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequencersList provides a mock function with given fields:
func (_m *MockLayerI) GetSequencersList() []*types.Sequencer {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSequencersList")
	}

	var r0 []*types.Sequencer
	if rf, ok := ret.Get(0).(func() []*types.Sequencer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Sequencer)
		}
	}

	return r0
}

// MockLayerI_GetSequencersList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequencersList'
type MockLayerI_GetSequencersList_Call struct {
	*mock.Call
}

// GetSequencersList is a helper method to define mock.On call
func (_e *MockLayerI_Expecter) GetSequencersList() *MockLayerI_GetSequencersList_Call {
	return &MockLayerI_GetSequencersList_Call{Call: _e.mock.On("GetSequencersList")}
}

func (_c *MockLayerI_GetSequencersList_Call) Run(run func()) *MockLayerI_GetSequencersList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLayerI_GetSequencersList_Call) Return(_a0 []*types.Sequencer) *MockLayerI_GetSequencersList_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_GetSequencersList_Call) RunAndReturn(run func() []*types.Sequencer) *MockLayerI_GetSequencersList_Call {
	_c.Call.Return(run)
	return _c
}

// Init provides a mock function with given fields: config, _a1, logger, options
func (_m *MockLayerI) Init(config settlement.Config, _a1 *pubsub.Server, logger types.Logger, options ...settlement.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, config, _a1, logger)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(settlement.Config, *pubsub.Server, types.Logger, ...settlement.Option) error); ok {
		r0 = rf(config, _a1, logger, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLayerI_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type MockLayerI_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
//   - config settlement.Config
//   - _a1 *pubsub.Server
//   - logger types.Logger
//   - options ...settlement.Option
func (_e *MockLayerI_Expecter) Init(config interface{}, _a1 interface{}, logger interface{}, options ...interface{}) *MockLayerI_Init_Call {
	return &MockLayerI_Init_Call{Call: _e.mock.On("Init",
		append([]interface{}{config, _a1, logger}, options...)...)}
}

func (_c *MockLayerI_Init_Call) Run(run func(config settlement.Config, _a1 *pubsub.Server, logger types.Logger, options ...settlement.Option)) *MockLayerI_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]settlement.Option, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(settlement.Option)
			}
		}
		run(args[0].(settlement.Config), args[1].(*pubsub.Server), args[2].(types.Logger), variadicArgs...)
	})
	return _c
}

func (_c *MockLayerI_Init_Call) Return(_a0 error) *MockLayerI_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_Init_Call) RunAndReturn(run func(settlement.Config, *pubsub.Server, types.Logger, ...settlement.Option) error) *MockLayerI_Init_Call {
	_c.Call.Return(run)
	return _c
}

// RetrieveBatch provides a mock function with given fields: stateIndex
func (_m *MockLayerI) RetrieveBatch(stateIndex ...uint64) (*settlement.ResultRetrieveBatch, error) {
	_va := make([]interface{}, len(stateIndex))
	for _i := range stateIndex {
		_va[_i] = stateIndex[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBatch")
	}

	var r0 *settlement.ResultRetrieveBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(...uint64) (*settlement.ResultRetrieveBatch, error)); ok {
		return rf(stateIndex...)
	}
	if rf, ok := ret.Get(0).(func(...uint64) *settlement.ResultRetrieveBatch); ok {
		r0 = rf(stateIndex...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultRetrieveBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(...uint64) error); ok {
		r1 = rf(stateIndex...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLayerI_RetrieveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RetrieveBatch'
type MockLayerI_RetrieveBatch_Call struct {
	*mock.Call
}

// RetrieveBatch is a helper method to define mock.On call
//   - stateIndex ...uint64
func (_e *MockLayerI_Expecter) RetrieveBatch(stateIndex ...interface{}) *MockLayerI_RetrieveBatch_Call {
	return &MockLayerI_RetrieveBatch_Call{Call: _e.mock.On("RetrieveBatch",
		append([]interface{}{}, stateIndex...)...)}
}

func (_c *MockLayerI_RetrieveBatch_Call) Run(run func(stateIndex ...uint64)) *MockLayerI_RetrieveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]uint64, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(uint64)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockLayerI_RetrieveBatch_Call) Return(_a0 *settlement.ResultRetrieveBatch, _a1 error) *MockLayerI_RetrieveBatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLayerI_RetrieveBatch_Call) RunAndReturn(run func(...uint64) (*settlement.ResultRetrieveBatch, error)) *MockLayerI_RetrieveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockLayerI) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLayerI_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockLayerI_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockLayerI_Expecter) Start() *MockLayerI_Start_Call {
	return &MockLayerI_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockLayerI_Start_Call) Run(run func()) *MockLayerI_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLayerI_Start_Call) Return(_a0 error) *MockLayerI_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_Start_Call) RunAndReturn(run func() error) *MockLayerI_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockLayerI) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLayerI_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockLayerI_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockLayerI_Expecter) Stop() *MockLayerI_Stop_Call {
	return &MockLayerI_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockLayerI_Stop_Call) Run(run func()) *MockLayerI_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockLayerI_Stop_Call) Return(_a0 error) *MockLayerI_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_Stop_Call) RunAndReturn(run func() error) *MockLayerI_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// SubmitBatch provides a mock function with given fields: batch, daClient, daResult
func (_m *MockLayerI) SubmitBatch(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch) error {
	ret := _m.Called(batch, daClient, daResult)

	if len(ret) == 0 {
		panic("no return value specified for SubmitBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Batch, da.Client, *da.ResultSubmitBatch) error); ok {
		r0 = rf(batch, daClient, daResult)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockLayerI_SubmitBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubmitBatch'
type MockLayerI_SubmitBatch_Call struct {
	*mock.Call
}

// SubmitBatch is a helper method to define mock.On call
//   - batch *types.Batch
//   - daClient da.Client
//   - daResult *da.ResultSubmitBatch
func (_e *MockLayerI_Expecter) SubmitBatch(batch interface{}, daClient interface{}, daResult interface{}) *MockLayerI_SubmitBatch_Call {
	return &MockLayerI_SubmitBatch_Call{Call: _e.mock.On("SubmitBatch", batch, daClient, daResult)}
}

func (_c *MockLayerI_SubmitBatch_Call) Run(run func(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch)) *MockLayerI_SubmitBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Batch), args[1].(da.Client), args[2].(*da.ResultSubmitBatch))
	})
	return _c
}

func (_c *MockLayerI_SubmitBatch_Call) Return(_a0 error) *MockLayerI_SubmitBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockLayerI_SubmitBatch_Call) RunAndReturn(run func(*types.Batch, da.Client, *da.ResultSubmitBatch) error) *MockLayerI_SubmitBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLayerI creates a new instance of MockLayerI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLayerI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLayerI {
	mock := &MockLayerI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
