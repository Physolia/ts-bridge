// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/ts-bridge/tsbridge (interfaces: SourceMetric)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/google/ts-bridge/storage"
	metric "google.golang.org/genproto/googleapis/api/metric"
	monitoring "google.golang.org/genproto/googleapis/monitoring/v3"
)

// MockSourceMetric is a mock of SourceMetric interface.
type MockSourceMetric struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMetricMockRecorder
}

// MockSourceMetricMockRecorder is the mock recorder for MockSourceMetric.
type MockSourceMetricMockRecorder struct {
	mock *MockSourceMetric
}

// NewMockSourceMetric creates a new mock instance.
func NewMockSourceMetric(ctrl *gomock.Controller) *MockSourceMetric {
	mock := &MockSourceMetric{ctrl: ctrl}
	mock.recorder = &MockSourceMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceMetric) EXPECT() *MockSourceMetricMockRecorder {
	return m.recorder
}

// Query mocks base method.
func (m *MockSourceMetric) Query() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query")
	ret0, _ := ret[0].(string)
	return ret0
}

// Query indicates an expected call of Query.
func (mr *MockSourceMetricMockRecorder) Query() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockSourceMetric)(nil).Query))
}

// StackdriverData mocks base method.
func (m *MockSourceMetric) StackdriverData(arg0 context.Context, arg1 time.Time, arg2 storage.MetricRecord) (*metric.MetricDescriptor, []*monitoring.TimeSeries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackdriverData", arg0, arg1, arg2)
	ret0, _ := ret[0].(*metric.MetricDescriptor)
	ret1, _ := ret[1].([]*monitoring.TimeSeries)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StackdriverData indicates an expected call of StackdriverData.
func (mr *MockSourceMetricMockRecorder) StackdriverData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackdriverData", reflect.TypeOf((*MockSourceMetric)(nil).StackdriverData), arg0, arg1, arg2)
}

// StackdriverName mocks base method.
func (m *MockSourceMetric) StackdriverName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackdriverName")
	ret0, _ := ret[0].(string)
	return ret0
}

// StackdriverName indicates an expected call of StackdriverName.
func (mr *MockSourceMetricMockRecorder) StackdriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackdriverName", reflect.TypeOf((*MockSourceMetric)(nil).StackdriverName))
}
