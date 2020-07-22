// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-mattermod/metrics (interfaces: Provider)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// IncreaseCronTaskErrors mocks base method
func (m *MockProvider) IncreaseCronTaskErrors(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncreaseCronTaskErrors", arg0)
}

// IncreaseCronTaskErrors indicates an expected call of IncreaseCronTaskErrors
func (mr *MockProviderMockRecorder) IncreaseCronTaskErrors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseCronTaskErrors", reflect.TypeOf((*MockProvider)(nil).IncreaseCronTaskErrors), arg0)
}

// IncreaseGithubCacheHits mocks base method
func (m *MockProvider) IncreaseGithubCacheHits(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncreaseGithubCacheHits", arg0, arg1)
}

// IncreaseGithubCacheHits indicates an expected call of IncreaseGithubCacheHits
func (mr *MockProviderMockRecorder) IncreaseGithubCacheHits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseGithubCacheHits", reflect.TypeOf((*MockProvider)(nil).IncreaseGithubCacheHits), arg0, arg1)
}

// IncreaseGithubCacheMisses mocks base method
func (m *MockProvider) IncreaseGithubCacheMisses(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncreaseGithubCacheMisses", arg0, arg1)
}

// IncreaseGithubCacheMisses indicates an expected call of IncreaseGithubCacheMisses
func (mr *MockProviderMockRecorder) IncreaseGithubCacheMisses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseGithubCacheMisses", reflect.TypeOf((*MockProvider)(nil).IncreaseGithubCacheMisses), arg0, arg1)
}

// IncreaseWebhookRequest mocks base method
func (m *MockProvider) IncreaseWebhookRequest(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncreaseWebhookRequest", arg0)
}

// IncreaseWebhookRequest indicates an expected call of IncreaseWebhookRequest
func (mr *MockProviderMockRecorder) IncreaseWebhookRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncreaseWebhookRequest", reflect.TypeOf((*MockProvider)(nil).IncreaseWebhookRequest), arg0)
}

// ObserveCronTaskDuration mocks base method
func (m *MockProvider) ObserveCronTaskDuration(arg0 string, arg1 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ObserveCronTaskDuration", arg0, arg1)
}

// ObserveCronTaskDuration indicates an expected call of ObserveCronTaskDuration
func (mr *MockProviderMockRecorder) ObserveCronTaskDuration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveCronTaskDuration", reflect.TypeOf((*MockProvider)(nil).ObserveCronTaskDuration), arg0, arg1)
}

// ObserveGithubRequestDuration mocks base method
func (m *MockProvider) ObserveGithubRequestDuration(arg0, arg1, arg2 string, arg3 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ObserveGithubRequestDuration", arg0, arg1, arg2, arg3)
}

// ObserveGithubRequestDuration indicates an expected call of ObserveGithubRequestDuration
func (mr *MockProviderMockRecorder) ObserveGithubRequestDuration(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveGithubRequestDuration", reflect.TypeOf((*MockProvider)(nil).ObserveGithubRequestDuration), arg0, arg1, arg2, arg3)
}

// ObserveHTTPRequestDuration mocks base method
func (m *MockProvider) ObserveHTTPRequestDuration(arg0, arg1, arg2 string, arg3 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ObserveHTTPRequestDuration", arg0, arg1, arg2, arg3)
}

// ObserveHTTPRequestDuration indicates an expected call of ObserveHTTPRequestDuration
func (mr *MockProviderMockRecorder) ObserveHTTPRequestDuration(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveHTTPRequestDuration", reflect.TypeOf((*MockProvider)(nil).ObserveHTTPRequestDuration), arg0, arg1, arg2, arg3)
}