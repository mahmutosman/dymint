// Code generated by mockery v2.43.0. DO NOT EDIT.

package settlement

import (
	da "github.com/dymensionxyz/dymint/da"
	mock "github.com/stretchr/testify/mock"

	settlement "github.com/dymensionxyz/dymint/settlement"

	types "github.com/dymensionxyz/dymint/types"
)

// MockHubClient is an autogenerated mock type for the HubClient type
type MockHubClient struct {
	mock.Mock
}

type MockHubClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHubClient) EXPECT() *MockHubClient_Expecter {
	return &MockHubClient_Expecter{mock: &_m.Mock}
}

// GetBatchAtIndex provides a mock function with given fields: rollappID, index
func (_m *MockHubClient) GetBatchAtIndex(rollappID string, index uint64) (*settlement.ResultRetrieveBatch, error) {
	ret := _m.Called(rollappID, index)

	if len(ret) == 0 {
		panic("no return value specified for GetBatchAtIndex")
	}

	var r0 *settlement.ResultRetrieveBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uint64) (*settlement.ResultRetrieveBatch, error)); ok {
		return rf(rollappID, index)
	}
	if rf, ok := ret.Get(0).(func(string, uint64) *settlement.ResultRetrieveBatch); ok {
		r0 = rf(rollappID, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultRetrieveBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(rollappID, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHubClient_GetBatchAtIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBatchAtIndex'
type MockHubClient_GetBatchAtIndex_Call struct {
	*mock.Call
}

// GetBatchAtIndex is a helper method to define mock.On call
//   - rollappID string
//   - index uint64
func (_e *MockHubClient_Expecter) GetBatchAtIndex(rollappID interface{}, index interface{}) *MockHubClient_GetBatchAtIndex_Call {
	return &MockHubClient_GetBatchAtIndex_Call{Call: _e.mock.On("GetBatchAtIndex", rollappID, index)}
}

func (_c *MockHubClient_GetBatchAtIndex_Call) Run(run func(rollappID string, index uint64)) *MockHubClient_GetBatchAtIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uint64))
	})
	return _c
}

func (_c *MockHubClient_GetBatchAtIndex_Call) Return(_a0 *settlement.ResultRetrieveBatch, _a1 error) *MockHubClient_GetBatchAtIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHubClient_GetBatchAtIndex_Call) RunAndReturn(run func(string, uint64) (*settlement.ResultRetrieveBatch, error)) *MockHubClient_GetBatchAtIndex_Call {
	_c.Call.Return(run)
	return _c
}

// GetHeightState provides a mock function with given fields: index
func (_m *MockHubClient) GetHeightState(index uint64) (*settlement.ResultGetHeightState, error) {
	ret := _m.Called(index)

	if len(ret) == 0 {
		panic("no return value specified for GetHeightState")
	}

	var r0 *settlement.ResultGetHeightState
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*settlement.ResultGetHeightState, error)); ok {
		return rf(index)
	}
	if rf, ok := ret.Get(0).(func(uint64) *settlement.ResultGetHeightState); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultGetHeightState)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHubClient_GetHeightState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeightState'
type MockHubClient_GetHeightState_Call struct {
	*mock.Call
}

// GetHeightState is a helper method to define mock.On call
//   - index uint64
func (_e *MockHubClient_Expecter) GetHeightState(index interface{}) *MockHubClient_GetHeightState_Call {
	return &MockHubClient_GetHeightState_Call{Call: _e.mock.On("GetHeightState", index)}
}

func (_c *MockHubClient_GetHeightState_Call) Run(run func(index uint64)) *MockHubClient_GetHeightState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64))
	})
	return _c
}

func (_c *MockHubClient_GetHeightState_Call) Return(_a0 *settlement.ResultGetHeightState, _a1 error) *MockHubClient_GetHeightState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHubClient_GetHeightState_Call) RunAndReturn(run func(uint64) (*settlement.ResultGetHeightState, error)) *MockHubClient_GetHeightState_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestBatch provides a mock function with given fields: rollappID
func (_m *MockHubClient) GetLatestBatch(rollappID string) (*settlement.ResultRetrieveBatch, error) {
	ret := _m.Called(rollappID)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBatch")
	}

	var r0 *settlement.ResultRetrieveBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*settlement.ResultRetrieveBatch, error)); ok {
		return rf(rollappID)
	}
	if rf, ok := ret.Get(0).(func(string) *settlement.ResultRetrieveBatch); ok {
		r0 = rf(rollappID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settlement.ResultRetrieveBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rollappID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHubClient_GetLatestBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestBatch'
type MockHubClient_GetLatestBatch_Call struct {
	*mock.Call
}

// GetLatestBatch is a helper method to define mock.On call
//   - rollappID string
func (_e *MockHubClient_Expecter) GetLatestBatch(rollappID interface{}) *MockHubClient_GetLatestBatch_Call {
	return &MockHubClient_GetLatestBatch_Call{Call: _e.mock.On("GetLatestBatch", rollappID)}
}

func (_c *MockHubClient_GetLatestBatch_Call) Run(run func(rollappID string)) *MockHubClient_GetLatestBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHubClient_GetLatestBatch_Call) Return(_a0 *settlement.ResultRetrieveBatch, _a1 error) *MockHubClient_GetLatestBatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHubClient_GetLatestBatch_Call) RunAndReturn(run func(string) (*settlement.ResultRetrieveBatch, error)) *MockHubClient_GetLatestBatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetSequencers provides a mock function with given fields: rollappID
func (_m *MockHubClient) GetSequencers(rollappID string) ([]*types.Sequencer, error) {
	ret := _m.Called(rollappID)

	if len(ret) == 0 {
		panic("no return value specified for GetSequencers")
	}

	var r0 []*types.Sequencer
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*types.Sequencer, error)); ok {
		return rf(rollappID)
	}
	if rf, ok := ret.Get(0).(func(string) []*types.Sequencer); ok {
		r0 = rf(rollappID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Sequencer)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rollappID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHubClient_GetSequencers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSequencers'
type MockHubClient_GetSequencers_Call struct {
	*mock.Call
}

// GetSequencers is a helper method to define mock.On call
//   - rollappID string
func (_e *MockHubClient_Expecter) GetSequencers(rollappID interface{}) *MockHubClient_GetSequencers_Call {
	return &MockHubClient_GetSequencers_Call{Call: _e.mock.On("GetSequencers", rollappID)}
}

func (_c *MockHubClient_GetSequencers_Call) Run(run func(rollappID string)) *MockHubClient_GetSequencers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHubClient_GetSequencers_Call) Return(_a0 []*types.Sequencer, _a1 error) *MockHubClient_GetSequencers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHubClient_GetSequencers_Call) RunAndReturn(run func(string) ([]*types.Sequencer, error)) *MockHubClient_GetSequencers_Call {
	_c.Call.Return(run)
	return _c
}

// PostBatch provides a mock function with given fields: batch, daClient, daResult
func (_m *MockHubClient) PostBatch(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch) error {
	ret := _m.Called(batch, daClient, daResult)

	if len(ret) == 0 {
		panic("no return value specified for PostBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Batch, da.Client, *da.ResultSubmitBatch) error); ok {
		r0 = rf(batch, daClient, daResult)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockHubClient_PostBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostBatch'
type MockHubClient_PostBatch_Call struct {
	*mock.Call
}

// PostBatch is a helper method to define mock.On call
//   - batch *types.Batch
//   - daClient da.Client
//   - daResult *da.ResultSubmitBatch
func (_e *MockHubClient_Expecter) PostBatch(batch interface{}, daClient interface{}, daResult interface{}) *MockHubClient_PostBatch_Call {
	return &MockHubClient_PostBatch_Call{Call: _e.mock.On("PostBatch", batch, daClient, daResult)}
}

func (_c *MockHubClient_PostBatch_Call) Run(run func(batch *types.Batch, daClient da.Client, daResult *da.ResultSubmitBatch)) *MockHubClient_PostBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Batch), args[1].(da.Client), args[2].(*da.ResultSubmitBatch))
	})
	return _c
}

func (_c *MockHubClient_PostBatch_Call) Return(_a0 error) *MockHubClient_PostBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHubClient_PostBatch_Call) RunAndReturn(run func(*types.Batch, da.Client, *da.ResultSubmitBatch) error) *MockHubClient_PostBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields:
func (_m *MockHubClient) Start() error {
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

// MockHubClient_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type MockHubClient_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
func (_e *MockHubClient_Expecter) Start() *MockHubClient_Start_Call {
	return &MockHubClient_Start_Call{Call: _e.mock.On("Start")}
}

func (_c *MockHubClient_Start_Call) Run(run func()) *MockHubClient_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHubClient_Start_Call) Return(_a0 error) *MockHubClient_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHubClient_Start_Call) RunAndReturn(run func() error) *MockHubClient_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *MockHubClient) Stop() error {
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

// MockHubClient_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type MockHubClient_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *MockHubClient_Expecter) Stop() *MockHubClient_Stop_Call {
	return &MockHubClient_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *MockHubClient_Stop_Call) Run(run func()) *MockHubClient_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHubClient_Stop_Call) Return(_a0 error) *MockHubClient_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHubClient_Stop_Call) RunAndReturn(run func() error) *MockHubClient_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHubClient creates a new instance of MockHubClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHubClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHubClient {
	mock := &MockHubClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}