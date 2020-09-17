// Code generated by MockGen. DO NOT EDIT.
// Source: queue.go

// Package server is a generated GoMock package.
package server

import (
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	headers "github.com/haraqa/haraqa/internal/headers"
)

// MockQueue is a mock of Queue interface
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// RootDir mocks base method
func (m *MockQueue) RootDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// RootDir indicates an expected call of RootDir
func (mr *MockQueueMockRecorder) RootDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootDir", reflect.TypeOf((*MockQueue)(nil).RootDir))
}

// Close mocks base method
func (m *MockQueue) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockQueueMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQueue)(nil).Close))
}

// ListTopics mocks base method
func (m *MockQueue) ListTopics(prefix, suffix, regex string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTopics", prefix, suffix, regex)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTopics indicates an expected call of ListTopics
func (mr *MockQueueMockRecorder) ListTopics(prefix, suffix, regex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockQueue)(nil).ListTopics), prefix, suffix, regex)
}

// CreateTopic mocks base method
func (m *MockQueue) CreateTopic(topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopic indicates an expected call of CreateTopic
func (mr *MockQueueMockRecorder) CreateTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockQueue)(nil).CreateTopic), topic)
}

// DeleteTopic mocks base method
func (m *MockQueue) DeleteTopic(topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTopic", topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTopic indicates an expected call of DeleteTopic
func (mr *MockQueueMockRecorder) DeleteTopic(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopic", reflect.TypeOf((*MockQueue)(nil).DeleteTopic), topic)
}

// ModifyTopic mocks base method
func (m *MockQueue) ModifyTopic(topic string, request headers.ModifyRequest) (*headers.TopicInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyTopic", topic, request)
	ret0, _ := ret[0].(*headers.TopicInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyTopic indicates an expected call of ModifyTopic
func (mr *MockQueueMockRecorder) ModifyTopic(topic, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyTopic", reflect.TypeOf((*MockQueue)(nil).ModifyTopic), topic, request)
}

// Produce mocks base method
func (m *MockQueue) Produce(topic string, msgSizes []int64, timestamp uint64, r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Produce", topic, msgSizes, timestamp, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce
func (mr *MockQueueMockRecorder) Produce(topic, msgSizes, timestamp, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockQueue)(nil).Produce), topic, msgSizes, timestamp, r)
}

// Consume mocks base method
func (m *MockQueue) Consume(topic string, id, limit int64, w http.ResponseWriter) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", topic, id, limit, w)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockQueueMockRecorder) Consume(topic, id, limit, w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockQueue)(nil).Consume), topic, id, limit, w)
}
