package census

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"github.com/insolar/insolar/network/consensus/common/endpoints"
	mm_census "github.com/insolar/insolar/network/consensus/gcpv2/api/census"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/profiles"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/proofs"
)

// MandateRegistryMock implements census.MandateRegistry
type MandateRegistryMock struct {
	t minimock.Tester

	funcFindRegisteredProfile          func(host endpoints.Inbound) (h1 profiles.Host)
	inspectFuncFindRegisteredProfile   func(host endpoints.Inbound)
	afterFindRegisteredProfileCounter  uint64
	beforeFindRegisteredProfileCounter uint64
	FindRegisteredProfileMock          mMandateRegistryMockFindRegisteredProfile

	funcGetCloudIdentity          func() (d1 cryptkit.DigestHolder)
	inspectFuncGetCloudIdentity   func()
	afterGetCloudIdentityCounter  uint64
	beforeGetCloudIdentityCounter uint64
	GetCloudIdentityMock          mMandateRegistryMockGetCloudIdentity

	funcGetConsensusConfiguration          func() (c1 mm_census.ConsensusConfiguration)
	inspectFuncGetConsensusConfiguration   func()
	afterGetConsensusConfigurationCounter  uint64
	beforeGetConsensusConfigurationCounter uint64
	GetConsensusConfigurationMock          mMandateRegistryMockGetConsensusConfiguration

	funcGetPrimingCloudHash          func() (c1 proofs.CloudStateHash)
	inspectFuncGetPrimingCloudHash   func()
	afterGetPrimingCloudHashCounter  uint64
	beforeGetPrimingCloudHashCounter uint64
	GetPrimingCloudHashMock          mMandateRegistryMockGetPrimingCloudHash
}

// NewMandateRegistryMock returns a mock for census.MandateRegistry
func NewMandateRegistryMock(t minimock.Tester) *MandateRegistryMock {
	m := &MandateRegistryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FindRegisteredProfileMock = mMandateRegistryMockFindRegisteredProfile{mock: m}
	m.FindRegisteredProfileMock.callArgs = []*MandateRegistryMockFindRegisteredProfileParams{}

	m.GetCloudIdentityMock = mMandateRegistryMockGetCloudIdentity{mock: m}

	m.GetConsensusConfigurationMock = mMandateRegistryMockGetConsensusConfiguration{mock: m}

	m.GetPrimingCloudHashMock = mMandateRegistryMockGetPrimingCloudHash{mock: m}

	return m
}

type mMandateRegistryMockFindRegisteredProfile struct {
	mock               *MandateRegistryMock
	defaultExpectation *MandateRegistryMockFindRegisteredProfileExpectation
	expectations       []*MandateRegistryMockFindRegisteredProfileExpectation

	callArgs []*MandateRegistryMockFindRegisteredProfileParams
	mutex    sync.RWMutex
}

// MandateRegistryMockFindRegisteredProfileExpectation specifies expectation struct of the MandateRegistry.FindRegisteredProfile
type MandateRegistryMockFindRegisteredProfileExpectation struct {
	mock    *MandateRegistryMock
	params  *MandateRegistryMockFindRegisteredProfileParams
	results *MandateRegistryMockFindRegisteredProfileResults
	Counter uint64
}

// MandateRegistryMockFindRegisteredProfileParams contains parameters of the MandateRegistry.FindRegisteredProfile
type MandateRegistryMockFindRegisteredProfileParams struct {
	host endpoints.Inbound
}

// MandateRegistryMockFindRegisteredProfileResults contains results of the MandateRegistry.FindRegisteredProfile
type MandateRegistryMockFindRegisteredProfileResults struct {
	h1 profiles.Host
}

// Expect sets up expected params for MandateRegistry.FindRegisteredProfile
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) Expect(host endpoints.Inbound) *mMandateRegistryMockFindRegisteredProfile {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("MandateRegistryMock.FindRegisteredProfile mock is already set by Set")
	}

	if mmFindRegisteredProfile.defaultExpectation == nil {
		mmFindRegisteredProfile.defaultExpectation = &MandateRegistryMockFindRegisteredProfileExpectation{}
	}

	mmFindRegisteredProfile.defaultExpectation.params = &MandateRegistryMockFindRegisteredProfileParams{host}
	for _, e := range mmFindRegisteredProfile.expectations {
		if minimock.Equal(e.params, mmFindRegisteredProfile.defaultExpectation.params) {
			mmFindRegisteredProfile.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFindRegisteredProfile.defaultExpectation.params)
		}
	}

	return mmFindRegisteredProfile
}

// Inspect accepts an inspector function that has same arguments as the MandateRegistry.FindRegisteredProfile
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) Inspect(f func(host endpoints.Inbound)) *mMandateRegistryMockFindRegisteredProfile {
	if mmFindRegisteredProfile.mock.inspectFuncFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("Inspect function is already set for MandateRegistryMock.FindRegisteredProfile")
	}

	mmFindRegisteredProfile.mock.inspectFuncFindRegisteredProfile = f

	return mmFindRegisteredProfile
}

// Return sets up results that will be returned by MandateRegistry.FindRegisteredProfile
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) Return(h1 profiles.Host) *MandateRegistryMock {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("MandateRegistryMock.FindRegisteredProfile mock is already set by Set")
	}

	if mmFindRegisteredProfile.defaultExpectation == nil {
		mmFindRegisteredProfile.defaultExpectation = &MandateRegistryMockFindRegisteredProfileExpectation{mock: mmFindRegisteredProfile.mock}
	}
	mmFindRegisteredProfile.defaultExpectation.results = &MandateRegistryMockFindRegisteredProfileResults{h1}
	return mmFindRegisteredProfile.mock
}

//Set uses given function f to mock the MandateRegistry.FindRegisteredProfile method
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) Set(f func(host endpoints.Inbound) (h1 profiles.Host)) *MandateRegistryMock {
	if mmFindRegisteredProfile.defaultExpectation != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("Default expectation is already set for the MandateRegistry.FindRegisteredProfile method")
	}

	if len(mmFindRegisteredProfile.expectations) > 0 {
		mmFindRegisteredProfile.mock.t.Fatalf("Some expectations are already set for the MandateRegistry.FindRegisteredProfile method")
	}

	mmFindRegisteredProfile.mock.funcFindRegisteredProfile = f
	return mmFindRegisteredProfile.mock
}

// When sets expectation for the MandateRegistry.FindRegisteredProfile which will trigger the result defined by the following
// Then helper
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) When(host endpoints.Inbound) *MandateRegistryMockFindRegisteredProfileExpectation {
	if mmFindRegisteredProfile.mock.funcFindRegisteredProfile != nil {
		mmFindRegisteredProfile.mock.t.Fatalf("MandateRegistryMock.FindRegisteredProfile mock is already set by Set")
	}

	expectation := &MandateRegistryMockFindRegisteredProfileExpectation{
		mock:   mmFindRegisteredProfile.mock,
		params: &MandateRegistryMockFindRegisteredProfileParams{host},
	}
	mmFindRegisteredProfile.expectations = append(mmFindRegisteredProfile.expectations, expectation)
	return expectation
}

// Then sets up MandateRegistry.FindRegisteredProfile return parameters for the expectation previously defined by the When method
func (e *MandateRegistryMockFindRegisteredProfileExpectation) Then(h1 profiles.Host) *MandateRegistryMock {
	e.results = &MandateRegistryMockFindRegisteredProfileResults{h1}
	return e.mock
}

// FindRegisteredProfile implements census.MandateRegistry
func (mmFindRegisteredProfile *MandateRegistryMock) FindRegisteredProfile(host endpoints.Inbound) (h1 profiles.Host) {
	mm_atomic.AddUint64(&mmFindRegisteredProfile.beforeFindRegisteredProfileCounter, 1)
	defer mm_atomic.AddUint64(&mmFindRegisteredProfile.afterFindRegisteredProfileCounter, 1)

	if mmFindRegisteredProfile.inspectFuncFindRegisteredProfile != nil {
		mmFindRegisteredProfile.inspectFuncFindRegisteredProfile(host)
	}

	params := &MandateRegistryMockFindRegisteredProfileParams{host}

	// Record call args
	mmFindRegisteredProfile.FindRegisteredProfileMock.mutex.Lock()
	mmFindRegisteredProfile.FindRegisteredProfileMock.callArgs = append(mmFindRegisteredProfile.FindRegisteredProfileMock.callArgs, params)
	mmFindRegisteredProfile.FindRegisteredProfileMock.mutex.Unlock()

	for _, e := range mmFindRegisteredProfile.FindRegisteredProfileMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.h1
		}
	}

	if mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.Counter, 1)
		want := mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.params
		got := MandateRegistryMockFindRegisteredProfileParams{host}
		if want != nil && !minimock.Equal(*want, got) {
			mmFindRegisteredProfile.t.Errorf("MandateRegistryMock.FindRegisteredProfile got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmFindRegisteredProfile.FindRegisteredProfileMock.defaultExpectation.results
		if results == nil {
			mmFindRegisteredProfile.t.Fatal("No results are set for the MandateRegistryMock.FindRegisteredProfile")
		}
		return (*results).h1
	}
	if mmFindRegisteredProfile.funcFindRegisteredProfile != nil {
		return mmFindRegisteredProfile.funcFindRegisteredProfile(host)
	}
	mmFindRegisteredProfile.t.Fatalf("Unexpected call to MandateRegistryMock.FindRegisteredProfile. %v", host)
	return
}

// FindRegisteredProfileAfterCounter returns a count of finished MandateRegistryMock.FindRegisteredProfile invocations
func (mmFindRegisteredProfile *MandateRegistryMock) FindRegisteredProfileAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindRegisteredProfile.afterFindRegisteredProfileCounter)
}

// FindRegisteredProfileBeforeCounter returns a count of MandateRegistryMock.FindRegisteredProfile invocations
func (mmFindRegisteredProfile *MandateRegistryMock) FindRegisteredProfileBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFindRegisteredProfile.beforeFindRegisteredProfileCounter)
}

// Calls returns a list of arguments used in each call to MandateRegistryMock.FindRegisteredProfile.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFindRegisteredProfile *mMandateRegistryMockFindRegisteredProfile) Calls() []*MandateRegistryMockFindRegisteredProfileParams {
	mmFindRegisteredProfile.mutex.RLock()

	argCopy := make([]*MandateRegistryMockFindRegisteredProfileParams, len(mmFindRegisteredProfile.callArgs))
	copy(argCopy, mmFindRegisteredProfile.callArgs)

	mmFindRegisteredProfile.mutex.RUnlock()

	return argCopy
}

// MinimockFindRegisteredProfileDone returns true if the count of the FindRegisteredProfile invocations corresponds
// the number of defined expectations
func (m *MandateRegistryMock) MinimockFindRegisteredProfileDone() bool {
	for _, e := range m.FindRegisteredProfileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindRegisteredProfileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindRegisteredProfile != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		return false
	}
	return true
}

// MinimockFindRegisteredProfileInspect logs each unmet expectation
func (m *MandateRegistryMock) MinimockFindRegisteredProfileInspect() {
	for _, e := range m.FindRegisteredProfileMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MandateRegistryMock.FindRegisteredProfile with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FindRegisteredProfileMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		if m.FindRegisteredProfileMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MandateRegistryMock.FindRegisteredProfile")
		} else {
			m.t.Errorf("Expected call to MandateRegistryMock.FindRegisteredProfile with params: %#v", *m.FindRegisteredProfileMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFindRegisteredProfile != nil && mm_atomic.LoadUint64(&m.afterFindRegisteredProfileCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.FindRegisteredProfile")
	}
}

type mMandateRegistryMockGetCloudIdentity struct {
	mock               *MandateRegistryMock
	defaultExpectation *MandateRegistryMockGetCloudIdentityExpectation
	expectations       []*MandateRegistryMockGetCloudIdentityExpectation
}

// MandateRegistryMockGetCloudIdentityExpectation specifies expectation struct of the MandateRegistry.GetCloudIdentity
type MandateRegistryMockGetCloudIdentityExpectation struct {
	mock *MandateRegistryMock

	results *MandateRegistryMockGetCloudIdentityResults
	Counter uint64
}

// MandateRegistryMockGetCloudIdentityResults contains results of the MandateRegistry.GetCloudIdentity
type MandateRegistryMockGetCloudIdentityResults struct {
	d1 cryptkit.DigestHolder
}

// Expect sets up expected params for MandateRegistry.GetCloudIdentity
func (mmGetCloudIdentity *mMandateRegistryMockGetCloudIdentity) Expect() *mMandateRegistryMockGetCloudIdentity {
	if mmGetCloudIdentity.mock.funcGetCloudIdentity != nil {
		mmGetCloudIdentity.mock.t.Fatalf("MandateRegistryMock.GetCloudIdentity mock is already set by Set")
	}

	if mmGetCloudIdentity.defaultExpectation == nil {
		mmGetCloudIdentity.defaultExpectation = &MandateRegistryMockGetCloudIdentityExpectation{}
	}

	return mmGetCloudIdentity
}

// Inspect accepts an inspector function that has same arguments as the MandateRegistry.GetCloudIdentity
func (mmGetCloudIdentity *mMandateRegistryMockGetCloudIdentity) Inspect(f func()) *mMandateRegistryMockGetCloudIdentity {
	if mmGetCloudIdentity.mock.inspectFuncGetCloudIdentity != nil {
		mmGetCloudIdentity.mock.t.Fatalf("Inspect function is already set for MandateRegistryMock.GetCloudIdentity")
	}

	mmGetCloudIdentity.mock.inspectFuncGetCloudIdentity = f

	return mmGetCloudIdentity
}

// Return sets up results that will be returned by MandateRegistry.GetCloudIdentity
func (mmGetCloudIdentity *mMandateRegistryMockGetCloudIdentity) Return(d1 cryptkit.DigestHolder) *MandateRegistryMock {
	if mmGetCloudIdentity.mock.funcGetCloudIdentity != nil {
		mmGetCloudIdentity.mock.t.Fatalf("MandateRegistryMock.GetCloudIdentity mock is already set by Set")
	}

	if mmGetCloudIdentity.defaultExpectation == nil {
		mmGetCloudIdentity.defaultExpectation = &MandateRegistryMockGetCloudIdentityExpectation{mock: mmGetCloudIdentity.mock}
	}
	mmGetCloudIdentity.defaultExpectation.results = &MandateRegistryMockGetCloudIdentityResults{d1}
	return mmGetCloudIdentity.mock
}

//Set uses given function f to mock the MandateRegistry.GetCloudIdentity method
func (mmGetCloudIdentity *mMandateRegistryMockGetCloudIdentity) Set(f func() (d1 cryptkit.DigestHolder)) *MandateRegistryMock {
	if mmGetCloudIdentity.defaultExpectation != nil {
		mmGetCloudIdentity.mock.t.Fatalf("Default expectation is already set for the MandateRegistry.GetCloudIdentity method")
	}

	if len(mmGetCloudIdentity.expectations) > 0 {
		mmGetCloudIdentity.mock.t.Fatalf("Some expectations are already set for the MandateRegistry.GetCloudIdentity method")
	}

	mmGetCloudIdentity.mock.funcGetCloudIdentity = f
	return mmGetCloudIdentity.mock
}

// GetCloudIdentity implements census.MandateRegistry
func (mmGetCloudIdentity *MandateRegistryMock) GetCloudIdentity() (d1 cryptkit.DigestHolder) {
	mm_atomic.AddUint64(&mmGetCloudIdentity.beforeGetCloudIdentityCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCloudIdentity.afterGetCloudIdentityCounter, 1)

	if mmGetCloudIdentity.inspectFuncGetCloudIdentity != nil {
		mmGetCloudIdentity.inspectFuncGetCloudIdentity()
	}

	if mmGetCloudIdentity.GetCloudIdentityMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCloudIdentity.GetCloudIdentityMock.defaultExpectation.Counter, 1)

		results := mmGetCloudIdentity.GetCloudIdentityMock.defaultExpectation.results
		if results == nil {
			mmGetCloudIdentity.t.Fatal("No results are set for the MandateRegistryMock.GetCloudIdentity")
		}
		return (*results).d1
	}
	if mmGetCloudIdentity.funcGetCloudIdentity != nil {
		return mmGetCloudIdentity.funcGetCloudIdentity()
	}
	mmGetCloudIdentity.t.Fatalf("Unexpected call to MandateRegistryMock.GetCloudIdentity.")
	return
}

// GetCloudIdentityAfterCounter returns a count of finished MandateRegistryMock.GetCloudIdentity invocations
func (mmGetCloudIdentity *MandateRegistryMock) GetCloudIdentityAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCloudIdentity.afterGetCloudIdentityCounter)
}

// GetCloudIdentityBeforeCounter returns a count of MandateRegistryMock.GetCloudIdentity invocations
func (mmGetCloudIdentity *MandateRegistryMock) GetCloudIdentityBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCloudIdentity.beforeGetCloudIdentityCounter)
}

// MinimockGetCloudIdentityDone returns true if the count of the GetCloudIdentity invocations corresponds
// the number of defined expectations
func (m *MandateRegistryMock) MinimockGetCloudIdentityDone() bool {
	for _, e := range m.GetCloudIdentityMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCloudIdentityMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCloudIdentityCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCloudIdentity != nil && mm_atomic.LoadUint64(&m.afterGetCloudIdentityCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCloudIdentityInspect logs each unmet expectation
func (m *MandateRegistryMock) MinimockGetCloudIdentityInspect() {
	for _, e := range m.GetCloudIdentityMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MandateRegistryMock.GetCloudIdentity")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCloudIdentityMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCloudIdentityCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetCloudIdentity")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCloudIdentity != nil && mm_atomic.LoadUint64(&m.afterGetCloudIdentityCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetCloudIdentity")
	}
}

type mMandateRegistryMockGetConsensusConfiguration struct {
	mock               *MandateRegistryMock
	defaultExpectation *MandateRegistryMockGetConsensusConfigurationExpectation
	expectations       []*MandateRegistryMockGetConsensusConfigurationExpectation
}

// MandateRegistryMockGetConsensusConfigurationExpectation specifies expectation struct of the MandateRegistry.GetConsensusConfiguration
type MandateRegistryMockGetConsensusConfigurationExpectation struct {
	mock *MandateRegistryMock

	results *MandateRegistryMockGetConsensusConfigurationResults
	Counter uint64
}

// MandateRegistryMockGetConsensusConfigurationResults contains results of the MandateRegistry.GetConsensusConfiguration
type MandateRegistryMockGetConsensusConfigurationResults struct {
	c1 mm_census.ConsensusConfiguration
}

// Expect sets up expected params for MandateRegistry.GetConsensusConfiguration
func (mmGetConsensusConfiguration *mMandateRegistryMockGetConsensusConfiguration) Expect() *mMandateRegistryMockGetConsensusConfiguration {
	if mmGetConsensusConfiguration.mock.funcGetConsensusConfiguration != nil {
		mmGetConsensusConfiguration.mock.t.Fatalf("MandateRegistryMock.GetConsensusConfiguration mock is already set by Set")
	}

	if mmGetConsensusConfiguration.defaultExpectation == nil {
		mmGetConsensusConfiguration.defaultExpectation = &MandateRegistryMockGetConsensusConfigurationExpectation{}
	}

	return mmGetConsensusConfiguration
}

// Inspect accepts an inspector function that has same arguments as the MandateRegistry.GetConsensusConfiguration
func (mmGetConsensusConfiguration *mMandateRegistryMockGetConsensusConfiguration) Inspect(f func()) *mMandateRegistryMockGetConsensusConfiguration {
	if mmGetConsensusConfiguration.mock.inspectFuncGetConsensusConfiguration != nil {
		mmGetConsensusConfiguration.mock.t.Fatalf("Inspect function is already set for MandateRegistryMock.GetConsensusConfiguration")
	}

	mmGetConsensusConfiguration.mock.inspectFuncGetConsensusConfiguration = f

	return mmGetConsensusConfiguration
}

// Return sets up results that will be returned by MandateRegistry.GetConsensusConfiguration
func (mmGetConsensusConfiguration *mMandateRegistryMockGetConsensusConfiguration) Return(c1 mm_census.ConsensusConfiguration) *MandateRegistryMock {
	if mmGetConsensusConfiguration.mock.funcGetConsensusConfiguration != nil {
		mmGetConsensusConfiguration.mock.t.Fatalf("MandateRegistryMock.GetConsensusConfiguration mock is already set by Set")
	}

	if mmGetConsensusConfiguration.defaultExpectation == nil {
		mmGetConsensusConfiguration.defaultExpectation = &MandateRegistryMockGetConsensusConfigurationExpectation{mock: mmGetConsensusConfiguration.mock}
	}
	mmGetConsensusConfiguration.defaultExpectation.results = &MandateRegistryMockGetConsensusConfigurationResults{c1}
	return mmGetConsensusConfiguration.mock
}

//Set uses given function f to mock the MandateRegistry.GetConsensusConfiguration method
func (mmGetConsensusConfiguration *mMandateRegistryMockGetConsensusConfiguration) Set(f func() (c1 mm_census.ConsensusConfiguration)) *MandateRegistryMock {
	if mmGetConsensusConfiguration.defaultExpectation != nil {
		mmGetConsensusConfiguration.mock.t.Fatalf("Default expectation is already set for the MandateRegistry.GetConsensusConfiguration method")
	}

	if len(mmGetConsensusConfiguration.expectations) > 0 {
		mmGetConsensusConfiguration.mock.t.Fatalf("Some expectations are already set for the MandateRegistry.GetConsensusConfiguration method")
	}

	mmGetConsensusConfiguration.mock.funcGetConsensusConfiguration = f
	return mmGetConsensusConfiguration.mock
}

// GetConsensusConfiguration implements census.MandateRegistry
func (mmGetConsensusConfiguration *MandateRegistryMock) GetConsensusConfiguration() (c1 mm_census.ConsensusConfiguration) {
	mm_atomic.AddUint64(&mmGetConsensusConfiguration.beforeGetConsensusConfigurationCounter, 1)
	defer mm_atomic.AddUint64(&mmGetConsensusConfiguration.afterGetConsensusConfigurationCounter, 1)

	if mmGetConsensusConfiguration.inspectFuncGetConsensusConfiguration != nil {
		mmGetConsensusConfiguration.inspectFuncGetConsensusConfiguration()
	}

	if mmGetConsensusConfiguration.GetConsensusConfigurationMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetConsensusConfiguration.GetConsensusConfigurationMock.defaultExpectation.Counter, 1)

		results := mmGetConsensusConfiguration.GetConsensusConfigurationMock.defaultExpectation.results
		if results == nil {
			mmGetConsensusConfiguration.t.Fatal("No results are set for the MandateRegistryMock.GetConsensusConfiguration")
		}
		return (*results).c1
	}
	if mmGetConsensusConfiguration.funcGetConsensusConfiguration != nil {
		return mmGetConsensusConfiguration.funcGetConsensusConfiguration()
	}
	mmGetConsensusConfiguration.t.Fatalf("Unexpected call to MandateRegistryMock.GetConsensusConfiguration.")
	return
}

// GetConsensusConfigurationAfterCounter returns a count of finished MandateRegistryMock.GetConsensusConfiguration invocations
func (mmGetConsensusConfiguration *MandateRegistryMock) GetConsensusConfigurationAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetConsensusConfiguration.afterGetConsensusConfigurationCounter)
}

// GetConsensusConfigurationBeforeCounter returns a count of MandateRegistryMock.GetConsensusConfiguration invocations
func (mmGetConsensusConfiguration *MandateRegistryMock) GetConsensusConfigurationBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetConsensusConfiguration.beforeGetConsensusConfigurationCounter)
}

// MinimockGetConsensusConfigurationDone returns true if the count of the GetConsensusConfiguration invocations corresponds
// the number of defined expectations
func (m *MandateRegistryMock) MinimockGetConsensusConfigurationDone() bool {
	for _, e := range m.GetConsensusConfigurationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetConsensusConfigurationMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetConsensusConfigurationCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetConsensusConfiguration != nil && mm_atomic.LoadUint64(&m.afterGetConsensusConfigurationCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetConsensusConfigurationInspect logs each unmet expectation
func (m *MandateRegistryMock) MinimockGetConsensusConfigurationInspect() {
	for _, e := range m.GetConsensusConfigurationMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MandateRegistryMock.GetConsensusConfiguration")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetConsensusConfigurationMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetConsensusConfigurationCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetConsensusConfiguration")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetConsensusConfiguration != nil && mm_atomic.LoadUint64(&m.afterGetConsensusConfigurationCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetConsensusConfiguration")
	}
}

type mMandateRegistryMockGetPrimingCloudHash struct {
	mock               *MandateRegistryMock
	defaultExpectation *MandateRegistryMockGetPrimingCloudHashExpectation
	expectations       []*MandateRegistryMockGetPrimingCloudHashExpectation
}

// MandateRegistryMockGetPrimingCloudHashExpectation specifies expectation struct of the MandateRegistry.GetPrimingCloudHash
type MandateRegistryMockGetPrimingCloudHashExpectation struct {
	mock *MandateRegistryMock

	results *MandateRegistryMockGetPrimingCloudHashResults
	Counter uint64
}

// MandateRegistryMockGetPrimingCloudHashResults contains results of the MandateRegistry.GetPrimingCloudHash
type MandateRegistryMockGetPrimingCloudHashResults struct {
	c1 proofs.CloudStateHash
}

// Expect sets up expected params for MandateRegistry.GetPrimingCloudHash
func (mmGetPrimingCloudHash *mMandateRegistryMockGetPrimingCloudHash) Expect() *mMandateRegistryMockGetPrimingCloudHash {
	if mmGetPrimingCloudHash.mock.funcGetPrimingCloudHash != nil {
		mmGetPrimingCloudHash.mock.t.Fatalf("MandateRegistryMock.GetPrimingCloudHash mock is already set by Set")
	}

	if mmGetPrimingCloudHash.defaultExpectation == nil {
		mmGetPrimingCloudHash.defaultExpectation = &MandateRegistryMockGetPrimingCloudHashExpectation{}
	}

	return mmGetPrimingCloudHash
}

// Inspect accepts an inspector function that has same arguments as the MandateRegistry.GetPrimingCloudHash
func (mmGetPrimingCloudHash *mMandateRegistryMockGetPrimingCloudHash) Inspect(f func()) *mMandateRegistryMockGetPrimingCloudHash {
	if mmGetPrimingCloudHash.mock.inspectFuncGetPrimingCloudHash != nil {
		mmGetPrimingCloudHash.mock.t.Fatalf("Inspect function is already set for MandateRegistryMock.GetPrimingCloudHash")
	}

	mmGetPrimingCloudHash.mock.inspectFuncGetPrimingCloudHash = f

	return mmGetPrimingCloudHash
}

// Return sets up results that will be returned by MandateRegistry.GetPrimingCloudHash
func (mmGetPrimingCloudHash *mMandateRegistryMockGetPrimingCloudHash) Return(c1 proofs.CloudStateHash) *MandateRegistryMock {
	if mmGetPrimingCloudHash.mock.funcGetPrimingCloudHash != nil {
		mmGetPrimingCloudHash.mock.t.Fatalf("MandateRegistryMock.GetPrimingCloudHash mock is already set by Set")
	}

	if mmGetPrimingCloudHash.defaultExpectation == nil {
		mmGetPrimingCloudHash.defaultExpectation = &MandateRegistryMockGetPrimingCloudHashExpectation{mock: mmGetPrimingCloudHash.mock}
	}
	mmGetPrimingCloudHash.defaultExpectation.results = &MandateRegistryMockGetPrimingCloudHashResults{c1}
	return mmGetPrimingCloudHash.mock
}

//Set uses given function f to mock the MandateRegistry.GetPrimingCloudHash method
func (mmGetPrimingCloudHash *mMandateRegistryMockGetPrimingCloudHash) Set(f func() (c1 proofs.CloudStateHash)) *MandateRegistryMock {
	if mmGetPrimingCloudHash.defaultExpectation != nil {
		mmGetPrimingCloudHash.mock.t.Fatalf("Default expectation is already set for the MandateRegistry.GetPrimingCloudHash method")
	}

	if len(mmGetPrimingCloudHash.expectations) > 0 {
		mmGetPrimingCloudHash.mock.t.Fatalf("Some expectations are already set for the MandateRegistry.GetPrimingCloudHash method")
	}

	mmGetPrimingCloudHash.mock.funcGetPrimingCloudHash = f
	return mmGetPrimingCloudHash.mock
}

// GetPrimingCloudHash implements census.MandateRegistry
func (mmGetPrimingCloudHash *MandateRegistryMock) GetPrimingCloudHash() (c1 proofs.CloudStateHash) {
	mm_atomic.AddUint64(&mmGetPrimingCloudHash.beforeGetPrimingCloudHashCounter, 1)
	defer mm_atomic.AddUint64(&mmGetPrimingCloudHash.afterGetPrimingCloudHashCounter, 1)

	if mmGetPrimingCloudHash.inspectFuncGetPrimingCloudHash != nil {
		mmGetPrimingCloudHash.inspectFuncGetPrimingCloudHash()
	}

	if mmGetPrimingCloudHash.GetPrimingCloudHashMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetPrimingCloudHash.GetPrimingCloudHashMock.defaultExpectation.Counter, 1)

		results := mmGetPrimingCloudHash.GetPrimingCloudHashMock.defaultExpectation.results
		if results == nil {
			mmGetPrimingCloudHash.t.Fatal("No results are set for the MandateRegistryMock.GetPrimingCloudHash")
		}
		return (*results).c1
	}
	if mmGetPrimingCloudHash.funcGetPrimingCloudHash != nil {
		return mmGetPrimingCloudHash.funcGetPrimingCloudHash()
	}
	mmGetPrimingCloudHash.t.Fatalf("Unexpected call to MandateRegistryMock.GetPrimingCloudHash.")
	return
}

// GetPrimingCloudHashAfterCounter returns a count of finished MandateRegistryMock.GetPrimingCloudHash invocations
func (mmGetPrimingCloudHash *MandateRegistryMock) GetPrimingCloudHashAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPrimingCloudHash.afterGetPrimingCloudHashCounter)
}

// GetPrimingCloudHashBeforeCounter returns a count of MandateRegistryMock.GetPrimingCloudHash invocations
func (mmGetPrimingCloudHash *MandateRegistryMock) GetPrimingCloudHashBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetPrimingCloudHash.beforeGetPrimingCloudHashCounter)
}

// MinimockGetPrimingCloudHashDone returns true if the count of the GetPrimingCloudHash invocations corresponds
// the number of defined expectations
func (m *MandateRegistryMock) MinimockGetPrimingCloudHashDone() bool {
	for _, e := range m.GetPrimingCloudHashMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPrimingCloudHashMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPrimingCloudHashCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPrimingCloudHash != nil && mm_atomic.LoadUint64(&m.afterGetPrimingCloudHashCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetPrimingCloudHashInspect logs each unmet expectation
func (m *MandateRegistryMock) MinimockGetPrimingCloudHashInspect() {
	for _, e := range m.GetPrimingCloudHashMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to MandateRegistryMock.GetPrimingCloudHash")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetPrimingCloudHashMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetPrimingCloudHashCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetPrimingCloudHash")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetPrimingCloudHash != nil && mm_atomic.LoadUint64(&m.afterGetPrimingCloudHashCounter) < 1 {
		m.t.Error("Expected call to MandateRegistryMock.GetPrimingCloudHash")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MandateRegistryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFindRegisteredProfileInspect()

		m.MinimockGetCloudIdentityInspect()

		m.MinimockGetConsensusConfigurationInspect()

		m.MinimockGetPrimingCloudHashInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MandateRegistryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MandateRegistryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFindRegisteredProfileDone() &&
		m.MinimockGetCloudIdentityDone() &&
		m.MinimockGetConsensusConfigurationDone() &&
		m.MinimockGetPrimingCloudHashDone()
}
