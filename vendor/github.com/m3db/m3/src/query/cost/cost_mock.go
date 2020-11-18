// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/cost (interfaces: ChainedEnforcer,ChainedReporter)

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package cost is a generated GoMock package.
package cost

import (
	"reflect"

	cost0 "github.com/m3db/m3/src/x/cost"

	"github.com/golang/mock/gomock"
)

// MockChainedEnforcer is a mock of ChainedEnforcer interface
type MockChainedEnforcer struct {
	ctrl     *gomock.Controller
	recorder *MockChainedEnforcerMockRecorder
}

// MockChainedEnforcerMockRecorder is the mock recorder for MockChainedEnforcer
type MockChainedEnforcerMockRecorder struct {
	mock *MockChainedEnforcer
}

// NewMockChainedEnforcer creates a new mock instance
func NewMockChainedEnforcer(ctrl *gomock.Controller) *MockChainedEnforcer {
	mock := &MockChainedEnforcer{ctrl: ctrl}
	mock.recorder = &MockChainedEnforcerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainedEnforcer) EXPECT() *MockChainedEnforcerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockChainedEnforcer) Add(arg0 cost0.Cost) cost0.Report {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(cost0.Report)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockChainedEnforcerMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockChainedEnforcer)(nil).Add), arg0)
}

// Child mocks base method
func (m *MockChainedEnforcer) Child(arg0 string) ChainedEnforcer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Child", arg0)
	ret0, _ := ret[0].(ChainedEnforcer)
	return ret0
}

// Child indicates an expected call of Child
func (mr *MockChainedEnforcerMockRecorder) Child(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Child", reflect.TypeOf((*MockChainedEnforcer)(nil).Child), arg0)
}

// Clone mocks base method
func (m *MockChainedEnforcer) Clone() cost0.Enforcer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(cost0.Enforcer)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockChainedEnforcerMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockChainedEnforcer)(nil).Clone))
}

// Close mocks base method
func (m *MockChainedEnforcer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockChainedEnforcerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChainedEnforcer)(nil).Close))
}

// Limit mocks base method
func (m *MockChainedEnforcer) Limit() cost0.Limit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Limit")
	ret0, _ := ret[0].(cost0.Limit)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockChainedEnforcerMockRecorder) Limit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockChainedEnforcer)(nil).Limit))
}

// Reporter mocks base method
func (m *MockChainedEnforcer) Reporter() cost0.EnforcerReporter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reporter")
	ret0, _ := ret[0].(cost0.EnforcerReporter)
	return ret0
}

// Reporter indicates an expected call of Reporter
func (mr *MockChainedEnforcerMockRecorder) Reporter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reporter", reflect.TypeOf((*MockChainedEnforcer)(nil).Reporter))
}

// State mocks base method
func (m *MockChainedEnforcer) State() (cost0.Report, cost0.Limit) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(cost0.Report)
	ret1, _ := ret[1].(cost0.Limit)
	return ret0, ret1
}

// State indicates an expected call of State
func (mr *MockChainedEnforcerMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockChainedEnforcer)(nil).State))
}

// MockChainedReporter is a mock of ChainedReporter interface
type MockChainedReporter struct {
	ctrl     *gomock.Controller
	recorder *MockChainedReporterMockRecorder
}

// MockChainedReporterMockRecorder is the mock recorder for MockChainedReporter
type MockChainedReporterMockRecorder struct {
	mock *MockChainedReporter
}

// NewMockChainedReporter creates a new mock instance
func NewMockChainedReporter(ctrl *gomock.Controller) *MockChainedReporter {
	mock := &MockChainedReporter{ctrl: ctrl}
	mock.recorder = &MockChainedReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChainedReporter) EXPECT() *MockChainedReporterMockRecorder {
	return m.recorder
}

// OnChildClose mocks base method
func (m *MockChainedReporter) OnChildClose(arg0 cost0.Cost) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnChildClose", arg0)
}

// OnChildClose indicates an expected call of OnChildClose
func (mr *MockChainedReporterMockRecorder) OnChildClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnChildClose", reflect.TypeOf((*MockChainedReporter)(nil).OnChildClose), arg0)
}

// OnClose mocks base method
func (m *MockChainedReporter) OnClose(arg0 cost0.Cost) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnClose", arg0)
}

// OnClose indicates an expected call of OnClose
func (mr *MockChainedReporterMockRecorder) OnClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnClose", reflect.TypeOf((*MockChainedReporter)(nil).OnClose), arg0)
}

// ReportCost mocks base method
func (m *MockChainedReporter) ReportCost(arg0 cost0.Cost) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportCost", arg0)
}

// ReportCost indicates an expected call of ReportCost
func (mr *MockChainedReporterMockRecorder) ReportCost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportCost", reflect.TypeOf((*MockChainedReporter)(nil).ReportCost), arg0)
}

// ReportCurrent mocks base method
func (m *MockChainedReporter) ReportCurrent(arg0 cost0.Cost) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportCurrent", arg0)
}

// ReportCurrent indicates an expected call of ReportCurrent
func (mr *MockChainedReporterMockRecorder) ReportCurrent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportCurrent", reflect.TypeOf((*MockChainedReporter)(nil).ReportCurrent), arg0)
}

// ReportOverLimit mocks base method
func (m *MockChainedReporter) ReportOverLimit(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportOverLimit", arg0)
}

// ReportOverLimit indicates an expected call of ReportOverLimit
func (mr *MockChainedReporterMockRecorder) ReportOverLimit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportOverLimit", reflect.TypeOf((*MockChainedReporter)(nil).ReportOverLimit), arg0)
}