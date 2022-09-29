// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorlib (interfaces: CIDataClient)

// Package jobruntestcaseanalyzer is a generated GoMock package.
package jobruntestcaseanalyzer

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	sets "k8s.io/apimachinery/pkg/util/sets"

	jobrunaggregatorapi "github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorapi"
	jobrunaggregatorlib "github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorlib"
)

// MockCIDataClient is a mock of CIDataClient interface
type MockCIDataClient struct {
	ctrl     *gomock.Controller
	recorder *MockCIDataClientMockRecorder
}

// MockCIDataClientMockRecorder is the mock recorder for MockCIDataClient
type MockCIDataClientMockRecorder struct {
	mock *MockCIDataClient
}

// NewMockCIDataClient creates a new mock instance
func NewMockCIDataClient(ctrl *gomock.Controller) *MockCIDataClient {
	mock := &MockCIDataClient{ctrl: ctrl}
	mock.recorder = &MockCIDataClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCIDataClient) EXPECT() *MockCIDataClientMockRecorder {
	return m.recorder
}

// GetBackendDisruptionStatisticsByJob mocks base method
func (m *MockCIDataClient) GetBackendDisruptionStatisticsByJob(arg0 context.Context, arg1 string) ([]jobrunaggregatorapi.BackendDisruptionStatisticsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendDisruptionStatisticsByJob", arg0, arg1)
	ret0, _ := ret[0].([]jobrunaggregatorapi.BackendDisruptionStatisticsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackendDisruptionStatisticsByJob indicates an expected call of GetBackendDisruptionStatisticsByJob
func (mr *MockCIDataClientMockRecorder) GetBackendDisruptionStatisticsByJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendDisruptionStatisticsByJob", reflect.TypeOf((*MockCIDataClient)(nil).GetBackendDisruptionStatisticsByJob), arg0, arg1)
}

// GetJobRunForJobNameAfterTime mocks base method
func (m *MockCIDataClient) GetJobRunForJobNameAfterTime(arg0 context.Context, arg1 string, arg2 time.Time) (*jobrunaggregatorapi.JobRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobRunForJobNameAfterTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorapi.JobRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobRunForJobNameAfterTime indicates an expected call of GetJobRunForJobNameAfterTime
func (mr *MockCIDataClientMockRecorder) GetJobRunForJobNameAfterTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobRunForJobNameAfterTime", reflect.TypeOf((*MockCIDataClient)(nil).GetJobRunForJobNameAfterTime), arg0, arg1, arg2)
}

// GetJobRunForJobNameBeforeTime mocks base method
func (m *MockCIDataClient) GetJobRunForJobNameBeforeTime(arg0 context.Context, arg1 string, arg2 time.Time) (*jobrunaggregatorapi.JobRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobRunForJobNameBeforeTime", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorapi.JobRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobRunForJobNameBeforeTime indicates an expected call of GetJobRunForJobNameBeforeTime
func (mr *MockCIDataClientMockRecorder) GetJobRunForJobNameBeforeTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobRunForJobNameBeforeTime", reflect.TypeOf((*MockCIDataClient)(nil).GetJobRunForJobNameBeforeTime), arg0, arg1, arg2)
}

// GetLastAggregationForJob mocks base method
func (m *MockCIDataClient) GetLastAggregationForJob(arg0 context.Context, arg1, arg2 string) (*jobrunaggregatorapi.AggregatedTestRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAggregationForJob", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorapi.AggregatedTestRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastAggregationForJob indicates an expected call of GetLastAggregationForJob
func (mr *MockCIDataClientMockRecorder) GetLastAggregationForJob(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAggregationForJob", reflect.TypeOf((*MockCIDataClient)(nil).GetLastAggregationForJob), arg0, arg1, arg2)
}

// GetLastJobRunWithAlertDataForJobName mocks base method
func (m *MockCIDataClient) GetLastJobRunWithAlertDataForJobName(arg0 context.Context, arg1 string) (*jobrunaggregatorapi.JobRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastJobRunWithAlertDataForJobName", arg0, arg1)
	ret0, _ := ret[0].(*jobrunaggregatorapi.JobRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastJobRunWithAlertDataForJobName indicates an expected call of GetLastJobRunWithAlertDataForJobName
func (mr *MockCIDataClientMockRecorder) GetLastJobRunWithAlertDataForJobName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastJobRunWithAlertDataForJobName", reflect.TypeOf((*MockCIDataClient)(nil).GetLastJobRunWithAlertDataForJobName), arg0, arg1)
}

// GetLastJobRunWithDisruptionDataForJobName mocks base method
func (m *MockCIDataClient) GetLastJobRunWithDisruptionDataForJobName(arg0 context.Context, arg1 string) (*jobrunaggregatorapi.JobRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastJobRunWithDisruptionDataForJobName", arg0, arg1)
	ret0, _ := ret[0].(*jobrunaggregatorapi.JobRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastJobRunWithDisruptionDataForJobName indicates an expected call of GetLastJobRunWithDisruptionDataForJobName
func (mr *MockCIDataClientMockRecorder) GetLastJobRunWithDisruptionDataForJobName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastJobRunWithDisruptionDataForJobName", reflect.TypeOf((*MockCIDataClient)(nil).GetLastJobRunWithDisruptionDataForJobName), arg0, arg1)
}

// GetLastJobRunWithTestRunDataForJobName mocks base method
func (m *MockCIDataClient) GetLastJobRunWithTestRunDataForJobName(arg0 context.Context, arg1 string) (*jobrunaggregatorapi.JobRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastJobRunWithTestRunDataForJobName", arg0, arg1)
	ret0, _ := ret[0].(*jobrunaggregatorapi.JobRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastJobRunWithTestRunDataForJobName indicates an expected call of GetLastJobRunWithTestRunDataForJobName
func (mr *MockCIDataClientMockRecorder) GetLastJobRunWithTestRunDataForJobName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastJobRunWithTestRunDataForJobName", reflect.TypeOf((*MockCIDataClient)(nil).GetLastJobRunWithTestRunDataForJobName), arg0, arg1)
}

// ListAggregatedTestRunsForJob mocks base method
func (m *MockCIDataClient) ListAggregatedTestRunsForJob(arg0 context.Context, arg1, arg2 string, arg3 time.Time) ([]jobrunaggregatorapi.AggregatedTestRunRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAggregatedTestRunsForJob", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]jobrunaggregatorapi.AggregatedTestRunRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAggregatedTestRunsForJob indicates an expected call of ListAggregatedTestRunsForJob
func (mr *MockCIDataClientMockRecorder) ListAggregatedTestRunsForJob(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAggregatedTestRunsForJob", reflect.TypeOf((*MockCIDataClient)(nil).ListAggregatedTestRunsForJob), arg0, arg1, arg2, arg3)
}

// ListAllJobs mocks base method
func (m *MockCIDataClient) ListAllJobs(arg0 context.Context) ([]jobrunaggregatorapi.JobRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllJobs", arg0)
	ret0, _ := ret[0].([]jobrunaggregatorapi.JobRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllJobs indicates an expected call of ListAllJobs
func (mr *MockCIDataClientMockRecorder) ListAllJobs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllJobs", reflect.TypeOf((*MockCIDataClient)(nil).ListAllJobs), arg0)
}

// ListReleaseTags mocks base method
func (m *MockCIDataClient) ListReleaseTags(arg0 context.Context) (sets.String, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleaseTags", arg0)
	ret0, _ := ret[0].(sets.String)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleaseTags indicates an expected call of ListReleaseTags
func (mr *MockCIDataClientMockRecorder) ListReleaseTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleaseTags", reflect.TypeOf((*MockCIDataClient)(nil).ListReleaseTags), arg0)
}

// ListUnifiedTestRunsForJobAfterDay mocks base method
func (m *MockCIDataClient) ListUnifiedTestRunsForJobAfterDay(arg0 context.Context, arg1 string, arg2 time.Time) (*jobrunaggregatorlib.UnifiedTestRunRowIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUnifiedTestRunsForJobAfterDay", arg0, arg1, arg2)
	ret0, _ := ret[0].(*jobrunaggregatorlib.UnifiedTestRunRowIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockCIDataClient) ListDisruptionHistoricalData(arg0 context.Context) ([]jobrunaggregatorapi.HistoricalData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDisruptionHistoricalData", arg0)
	ret0, _ := ret[0].([]*jobrunaggregatorapi.DisruptionHistoricalDataRow)
	ret1, _ := ret[1].(error)
	return jobrunaggregatorapi.ConvertToHistoricalData(ret0), ret1
}
func (m *MockCIDataClient) ListAlertHistoricalData(arg0 context.Context) ([]jobrunaggregatorapi.HistoricalData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAlertHistoricalData", arg0)
	ret0, _ := ret[0].([]*jobrunaggregatorapi.AlertHistoricalDataRow)
	ret1, _ := ret[1].(error)
	return jobrunaggregatorapi.ConvertToHistoricalData(ret0), ret1
}

// ListUnifiedTestRunsForJobAfterDay indicates an expected call of ListUnifiedTestRunsForJobAfterDay
func (mr *MockCIDataClientMockRecorder) ListUnifiedTestRunsForJobAfterDay(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnifiedTestRunsForJobAfterDay", reflect.TypeOf((*MockCIDataClient)(nil).ListUnifiedTestRunsForJobAfterDay), arg0, arg1, arg2)
}

func (mr *MockCIDataClientMockRecorder) ListDisruptionHistoricalData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDisruptionHistoricalData", reflect.TypeOf((*MockCIDataClient)(nil).ListDisruptionHistoricalData), arg0)
}

func (mr *MockCIDataClientMockRecorder) ListAlertHistoricalData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUnifiedTestRunsForJobAfterDay", reflect.TypeOf((*MockCIDataClient)(nil).ListAlertHistoricalData), arg0)
}
