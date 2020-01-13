// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	haraqa "github.com/haraqa/haraqa"
	io "io"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CreateTopic mocks base method
func (m *MockClient) CreateTopic(ctx context.Context, topic []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopic", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTopic indicates an expected call of CreateTopic
func (mr *MockClientMockRecorder) CreateTopic(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopic", reflect.TypeOf((*MockClient)(nil).CreateTopic), ctx, topic)
}

// DeleteTopic mocks base method
func (m *MockClient) DeleteTopic(ctx context.Context, topic []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTopic", ctx, topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTopic indicates an expected call of DeleteTopic
func (mr *MockClientMockRecorder) DeleteTopic(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTopic", reflect.TypeOf((*MockClient)(nil).DeleteTopic), ctx, topic)
}

// ListTopics mocks base method
func (m *MockClient) ListTopics(ctx context.Context, prefix, suffix, regex string) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTopics", ctx, prefix, suffix, regex)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTopics indicates an expected call of ListTopics
func (mr *MockClientMockRecorder) ListTopics(ctx, prefix, suffix, regex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockClient)(nil).ListTopics), ctx, prefix, suffix, regex)
}

// WatchTopics mocks base method
func (m *MockClient) WatchTopics(ctx context.Context, ch chan haraqa.WatchEvent, topics ...[]byte) (io.Closer, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ch}
	for _, a := range topics {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchTopics", varargs...)
	ret0, _ := ret[0].(io.Closer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchTopics indicates an expected call of WatchTopics
func (mr *MockClientMockRecorder) WatchTopics(ctx, ch interface{}, topics ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ch}, topics...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchTopics", reflect.TypeOf((*MockClient)(nil).WatchTopics), varargs...)
}

// Offsets mocks base method
func (m *MockClient) Offsets(ctx context.Context, topic []byte) (int64, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offsets", ctx, topic)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Offsets indicates an expected call of Offsets
func (mr *MockClientMockRecorder) Offsets(ctx, topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offsets", reflect.TypeOf((*MockClient)(nil).Offsets), ctx, topic)
}

// Produce mocks base method
func (m *MockClient) Produce(ctx context.Context, topic []byte, msgs ...[]byte) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, topic}
	for _, a := range msgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Produce", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce
func (mr *MockClientMockRecorder) Produce(ctx, topic interface{}, msgs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, topic}, msgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockClient)(nil).Produce), varargs...)
}

// ProduceLoop mocks base method
func (m *MockClient) ProduceLoop(ctx context.Context, topic []byte, ch chan haraqa.ProduceMsg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceLoop", ctx, topic, ch)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceLoop indicates an expected call of ProduceLoop
func (mr *MockClientMockRecorder) ProduceLoop(ctx, topic, ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceLoop", reflect.TypeOf((*MockClient)(nil).ProduceLoop), ctx, topic, ch)
}

// Consume mocks base method
func (m *MockClient) Consume(ctx context.Context, topic []byte, offset, maxBatchSize int64, buf *haraqa.ConsumeBuffer) ([][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", ctx, topic, offset, maxBatchSize, buf)
	ret0, _ := ret[0].([][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockClientMockRecorder) Consume(ctx, topic, offset, maxBatchSize, buf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockClient)(nil).Consume), ctx, topic, offset, maxBatchSize, buf)
}

// Close mocks base method
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}
