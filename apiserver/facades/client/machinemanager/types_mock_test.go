// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/client/machinemanager (interfaces: Machine,Application,Unit,Charm,CharmMeta)

// Package machinemanager is a generated GoMock package.
package machinemanager

import (
	gomock "github.com/golang/mock/gomock"
	v8 "github.com/juju/charm/v8"
	model "github.com/juju/juju/core/model"
	status "github.com/juju/juju/core/status"
	state "github.com/juju/juju/state"
	v4 "github.com/juju/names/v4"
	reflect "reflect"
	time "time"
)

// MockMachine is a mock of Machine interface
type MockMachine struct {
	ctrl     *gomock.Controller
	recorder *MockMachineMockRecorder
}

// MockMachineMockRecorder is the mock recorder for MockMachine
type MockMachineMockRecorder struct {
	mock *MockMachine
}

// NewMockMachine creates a new mock instance
func NewMockMachine(ctrl *gomock.Controller) *MockMachine {
	mock := &MockMachine{ctrl: ctrl}
	mock.recorder = &MockMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine) EXPECT() *MockMachineMockRecorder {
	return m.recorder
}

// ApplicationNames mocks base method
func (m *MockMachine) ApplicationNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationNames indicates an expected call of ApplicationNames
func (mr *MockMachineMockRecorder) ApplicationNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationNames", reflect.TypeOf((*MockMachine)(nil).ApplicationNames))
}

// CompleteUpgradeSeries mocks base method
func (m *MockMachine) CompleteUpgradeSeries() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteUpgradeSeries")
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteUpgradeSeries indicates an expected call of CompleteUpgradeSeries
func (mr *MockMachineMockRecorder) CompleteUpgradeSeries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteUpgradeSeries", reflect.TypeOf((*MockMachine)(nil).CompleteUpgradeSeries))
}

// CreateUpgradeSeriesLock mocks base method
func (m *MockMachine) CreateUpgradeSeriesLock(arg0 []string, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpgradeSeriesLock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUpgradeSeriesLock indicates an expected call of CreateUpgradeSeriesLock
func (mr *MockMachineMockRecorder) CreateUpgradeSeriesLock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpgradeSeriesLock", reflect.TypeOf((*MockMachine)(nil).CreateUpgradeSeriesLock), arg0, arg1)
}

// Destroy mocks base method
func (m *MockMachine) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockMachineMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockMachine)(nil).Destroy))
}

// ForceDestroy mocks base method
func (m *MockMachine) ForceDestroy(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceDestroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceDestroy indicates an expected call of ForceDestroy
func (mr *MockMachineMockRecorder) ForceDestroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceDestroy", reflect.TypeOf((*MockMachine)(nil).ForceDestroy), arg0)
}

// GetUpgradeSeriesMessages mocks base method
func (m *MockMachine) GetUpgradeSeriesMessages() ([]string, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpgradeSeriesMessages")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUpgradeSeriesMessages indicates an expected call of GetUpgradeSeriesMessages
func (mr *MockMachineMockRecorder) GetUpgradeSeriesMessages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpgradeSeriesMessages", reflect.TypeOf((*MockMachine)(nil).GetUpgradeSeriesMessages))
}

// Id mocks base method
func (m *MockMachine) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id
func (mr *MockMachineMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockMachine)(nil).Id))
}

// IsLockedForSeriesUpgrade mocks base method
func (m *MockMachine) IsLockedForSeriesUpgrade() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLockedForSeriesUpgrade")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsLockedForSeriesUpgrade indicates an expected call of IsLockedForSeriesUpgrade
func (mr *MockMachineMockRecorder) IsLockedForSeriesUpgrade() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLockedForSeriesUpgrade", reflect.TypeOf((*MockMachine)(nil).IsLockedForSeriesUpgrade))
}

// IsManager mocks base method
func (m *MockMachine) IsManager() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsManager")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsManager indicates an expected call of IsManager
func (mr *MockMachineMockRecorder) IsManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsManager", reflect.TypeOf((*MockMachine)(nil).IsManager))
}

// Principals mocks base method
func (m *MockMachine) Principals() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Principals")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Principals indicates an expected call of Principals
func (mr *MockMachineMockRecorder) Principals() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Principals", reflect.TypeOf((*MockMachine)(nil).Principals))
}

// RemoveUpgradeSeriesLock mocks base method
func (m *MockMachine) RemoveUpgradeSeriesLock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUpgradeSeriesLock")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUpgradeSeriesLock indicates an expected call of RemoveUpgradeSeriesLock
func (mr *MockMachineMockRecorder) RemoveUpgradeSeriesLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUpgradeSeriesLock", reflect.TypeOf((*MockMachine)(nil).RemoveUpgradeSeriesLock))
}

// Series mocks base method
func (m *MockMachine) Series() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Series")
	ret0, _ := ret[0].(string)
	return ret0
}

// Series indicates an expected call of Series
func (mr *MockMachineMockRecorder) Series() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Series", reflect.TypeOf((*MockMachine)(nil).Series))
}

// SetKeepInstance mocks base method
func (m *MockMachine) SetKeepInstance(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetKeepInstance", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetKeepInstance indicates an expected call of SetKeepInstance
func (mr *MockMachineMockRecorder) SetKeepInstance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKeepInstance", reflect.TypeOf((*MockMachine)(nil).SetKeepInstance), arg0)
}

// Units mocks base method
func (m *MockMachine) Units() ([]Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockMachineMockRecorder) Units() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockMachine)(nil).Units))
}

// UpgradeSeriesStatus mocks base method
func (m *MockMachine) UpgradeSeriesStatus() (model.UpgradeSeriesStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeSeriesStatus")
	ret0, _ := ret[0].(model.UpgradeSeriesStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeSeriesStatus indicates an expected call of UpgradeSeriesStatus
func (mr *MockMachineMockRecorder) UpgradeSeriesStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeSeriesStatus", reflect.TypeOf((*MockMachine)(nil).UpgradeSeriesStatus))
}

// VerifyUnitsSeries mocks base method
func (m *MockMachine) VerifyUnitsSeries(arg0 []string, arg1 string, arg2 bool) ([]Unit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUnitsSeries", arg0, arg1, arg2)
	ret0, _ := ret[0].([]Unit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUnitsSeries indicates an expected call of VerifyUnitsSeries
func (mr *MockMachineMockRecorder) VerifyUnitsSeries(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUnitsSeries", reflect.TypeOf((*MockMachine)(nil).VerifyUnitsSeries), arg0, arg1, arg2)
}

// WatchUpgradeSeriesNotifications mocks base method
func (m *MockMachine) WatchUpgradeSeriesNotifications() (state.NotifyWatcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpgradeSeriesNotifications")
	ret0, _ := ret[0].(state.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpgradeSeriesNotifications indicates an expected call of WatchUpgradeSeriesNotifications
func (mr *MockMachineMockRecorder) WatchUpgradeSeriesNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpgradeSeriesNotifications", reflect.TypeOf((*MockMachine)(nil).WatchUpgradeSeriesNotifications))
}

// MockApplication is a mock of Application interface
type MockApplication struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationMockRecorder
}

// MockApplicationMockRecorder is the mock recorder for MockApplication
type MockApplicationMockRecorder struct {
	mock *MockApplication
}

// NewMockApplication creates a new mock instance
func NewMockApplication(ctrl *gomock.Controller) *MockApplication {
	mock := &MockApplication{ctrl: ctrl}
	mock.recorder = &MockApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplication) EXPECT() *MockApplicationMockRecorder {
	return m.recorder
}

// Charm mocks base method
func (m *MockApplication) Charm() (Charm, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charm")
	ret0, _ := ret[0].(Charm)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Charm indicates an expected call of Charm
func (mr *MockApplicationMockRecorder) Charm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockApplication)(nil).Charm))
}

// MockUnit is a mock of Unit interface
type MockUnit struct {
	ctrl     *gomock.Controller
	recorder *MockUnitMockRecorder
}

// MockUnitMockRecorder is the mock recorder for MockUnit
type MockUnitMockRecorder struct {
	mock *MockUnit
}

// NewMockUnit creates a new mock instance
func NewMockUnit(ctrl *gomock.Controller) *MockUnit {
	mock := &MockUnit{ctrl: ctrl}
	mock.recorder = &MockUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnit) EXPECT() *MockUnitMockRecorder {
	return m.recorder
}

// AgentStatus mocks base method
func (m *MockUnit) AgentStatus() (status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentStatus")
	ret0, _ := ret[0].(status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AgentStatus indicates an expected call of AgentStatus
func (mr *MockUnitMockRecorder) AgentStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentStatus", reflect.TypeOf((*MockUnit)(nil).AgentStatus))
}

// Name mocks base method
func (m *MockUnit) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockUnitMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockUnit)(nil).Name))
}

// Status mocks base method
func (m *MockUnit) Status() (status.StatusInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(status.StatusInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockUnitMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockUnit)(nil).Status))
}

// UnitTag mocks base method
func (m *MockUnit) UnitTag() v4.UnitTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnitTag")
	ret0, _ := ret[0].(v4.UnitTag)
	return ret0
}

// UnitTag indicates an expected call of UnitTag
func (mr *MockUnitMockRecorder) UnitTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnitTag", reflect.TypeOf((*MockUnit)(nil).UnitTag))
}

// MockCharm is a mock of Charm interface
type MockCharm struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMockRecorder
}

// MockCharmMockRecorder is the mock recorder for MockCharm
type MockCharmMockRecorder struct {
	mock *MockCharm
}

// NewMockCharm creates a new mock instance
func NewMockCharm(ctrl *gomock.Controller) *MockCharm {
	mock := &MockCharm{ctrl: ctrl}
	mock.recorder = &MockCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharm) EXPECT() *MockCharmMockRecorder {
	return m.recorder
}

// Meta mocks base method
func (m *MockCharm) Meta() CharmMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(CharmMeta)
	return ret0
}

// Meta indicates an expected call of Meta
func (mr *MockCharmMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharm)(nil).Meta))
}

// String mocks base method
func (m *MockCharm) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockCharmMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockCharm)(nil).String))
}

// URL mocks base method
func (m *MockCharm) URL() *v8.URL {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(*v8.URL)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockCharmMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockCharm)(nil).URL))
}

// MockCharmMeta is a mock of CharmMeta interface
type MockCharmMeta struct {
	ctrl     *gomock.Controller
	recorder *MockCharmMetaMockRecorder
}

// MockCharmMetaMockRecorder is the mock recorder for MockCharmMeta
type MockCharmMetaMockRecorder struct {
	mock *MockCharmMeta
}

// NewMockCharmMeta creates a new mock instance
func NewMockCharmMeta(ctrl *gomock.Controller) *MockCharmMeta {
	mock := &MockCharmMeta{ctrl: ctrl}
	mock.recorder = &MockCharmMetaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharmMeta) EXPECT() *MockCharmMetaMockRecorder {
	return m.recorder
}

// ComputedSeries mocks base method
func (m *MockCharmMeta) ComputedSeries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputedSeries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ComputedSeries indicates an expected call of ComputedSeries
func (mr *MockCharmMetaMockRecorder) ComputedSeries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputedSeries", reflect.TypeOf((*MockCharmMeta)(nil).ComputedSeries))
}
