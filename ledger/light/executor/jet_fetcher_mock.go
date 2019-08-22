package executor

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// JetFetcherMock implements executor.JetFetcher
type JetFetcherMock struct {
	t minimock.Tester

	funcFetch          func(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber) (ip1 *insolar.ID, err error)
	inspectFuncFetch   func(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber)
	afterFetchCounter  uint64
	beforeFetchCounter uint64
	FetchMock          mJetFetcherMockFetch

	funcRelease          func(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber)
	inspectFuncRelease   func(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber)
	afterReleaseCounter  uint64
	beforeReleaseCounter uint64
	ReleaseMock          mJetFetcherMockRelease
}

// NewJetFetcherMock returns a mock for executor.JetFetcher
func NewJetFetcherMock(t minimock.Tester) *JetFetcherMock {
	m := &JetFetcherMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FetchMock = mJetFetcherMockFetch{mock: m}
	m.FetchMock.callArgs = []*JetFetcherMockFetchParams{}

	m.ReleaseMock = mJetFetcherMockRelease{mock: m}
	m.ReleaseMock.callArgs = []*JetFetcherMockReleaseParams{}

	return m
}

type mJetFetcherMockFetch struct {
	mock               *JetFetcherMock
	defaultExpectation *JetFetcherMockFetchExpectation
	expectations       []*JetFetcherMockFetchExpectation

	callArgs []*JetFetcherMockFetchParams
	mutex    sync.RWMutex
}

// JetFetcherMockFetchExpectation specifies expectation struct of the JetFetcher.Fetch
type JetFetcherMockFetchExpectation struct {
	mock    *JetFetcherMock
	params  *JetFetcherMockFetchParams
	results *JetFetcherMockFetchResults
	Counter uint64
}

// JetFetcherMockFetchParams contains parameters of the JetFetcher.Fetch
type JetFetcherMockFetchParams struct {
	ctx    context.Context
	target insolar.ID
	pulse  insolar.PulseNumber
}

// JetFetcherMockFetchResults contains results of the JetFetcher.Fetch
type JetFetcherMockFetchResults struct {
	ip1 *insolar.ID
	err error
}

// Expect sets up expected params for JetFetcher.Fetch
func (mmFetch *mJetFetcherMockFetch) Expect(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber) *mJetFetcherMockFetch {
	if mmFetch.mock.funcFetch != nil {
		mmFetch.mock.t.Fatalf("JetFetcherMock.Fetch mock is already set by Set")
	}

	if mmFetch.defaultExpectation == nil {
		mmFetch.defaultExpectation = &JetFetcherMockFetchExpectation{}
	}

	mmFetch.defaultExpectation.params = &JetFetcherMockFetchParams{ctx, target, pulse}
	for _, e := range mmFetch.expectations {
		if minimock.Equal(e.params, mmFetch.defaultExpectation.params) {
			mmFetch.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFetch.defaultExpectation.params)
		}
	}

	return mmFetch
}

// Inspect accepts an inspector function that has same arguments as the JetFetcher.Fetch
func (mmFetch *mJetFetcherMockFetch) Inspect(f func(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber)) *mJetFetcherMockFetch {
	if mmFetch.mock.inspectFuncFetch != nil {
		mmFetch.mock.t.Fatalf("Inspect function is already set for JetFetcherMock.Fetch")
	}

	mmFetch.mock.inspectFuncFetch = f

	return mmFetch
}

// Return sets up results that will be returned by JetFetcher.Fetch
func (mmFetch *mJetFetcherMockFetch) Return(ip1 *insolar.ID, err error) *JetFetcherMock {
	if mmFetch.mock.funcFetch != nil {
		mmFetch.mock.t.Fatalf("JetFetcherMock.Fetch mock is already set by Set")
	}

	if mmFetch.defaultExpectation == nil {
		mmFetch.defaultExpectation = &JetFetcherMockFetchExpectation{mock: mmFetch.mock}
	}
	mmFetch.defaultExpectation.results = &JetFetcherMockFetchResults{ip1, err}
	return mmFetch.mock
}

//Set uses given function f to mock the JetFetcher.Fetch method
func (mmFetch *mJetFetcherMockFetch) Set(f func(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber) (ip1 *insolar.ID, err error)) *JetFetcherMock {
	if mmFetch.defaultExpectation != nil {
		mmFetch.mock.t.Fatalf("Default expectation is already set for the JetFetcher.Fetch method")
	}

	if len(mmFetch.expectations) > 0 {
		mmFetch.mock.t.Fatalf("Some expectations are already set for the JetFetcher.Fetch method")
	}

	mmFetch.mock.funcFetch = f
	return mmFetch.mock
}

// When sets expectation for the JetFetcher.Fetch which will trigger the result defined by the following
// Then helper
func (mmFetch *mJetFetcherMockFetch) When(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber) *JetFetcherMockFetchExpectation {
	if mmFetch.mock.funcFetch != nil {
		mmFetch.mock.t.Fatalf("JetFetcherMock.Fetch mock is already set by Set")
	}

	expectation := &JetFetcherMockFetchExpectation{
		mock:   mmFetch.mock,
		params: &JetFetcherMockFetchParams{ctx, target, pulse},
	}
	mmFetch.expectations = append(mmFetch.expectations, expectation)
	return expectation
}

// Then sets up JetFetcher.Fetch return parameters for the expectation previously defined by the When method
func (e *JetFetcherMockFetchExpectation) Then(ip1 *insolar.ID, err error) *JetFetcherMock {
	e.results = &JetFetcherMockFetchResults{ip1, err}
	return e.mock
}

// Fetch implements executor.JetFetcher
func (mmFetch *JetFetcherMock) Fetch(ctx context.Context, target insolar.ID, pulse insolar.PulseNumber) (ip1 *insolar.ID, err error) {
	mm_atomic.AddUint64(&mmFetch.beforeFetchCounter, 1)
	defer mm_atomic.AddUint64(&mmFetch.afterFetchCounter, 1)

	if mmFetch.inspectFuncFetch != nil {
		mmFetch.inspectFuncFetch(ctx, target, pulse)
	}

	params := &JetFetcherMockFetchParams{ctx, target, pulse}

	// Record call args
	mmFetch.FetchMock.mutex.Lock()
	mmFetch.FetchMock.callArgs = append(mmFetch.FetchMock.callArgs, params)
	mmFetch.FetchMock.mutex.Unlock()

	for _, e := range mmFetch.FetchMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ip1, e.results.err
		}
	}

	if mmFetch.FetchMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFetch.FetchMock.defaultExpectation.Counter, 1)
		want := mmFetch.FetchMock.defaultExpectation.params
		got := JetFetcherMockFetchParams{ctx, target, pulse}
		if want != nil && !minimock.Equal(*want, got) {
			mmFetch.t.Errorf("JetFetcherMock.Fetch got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmFetch.FetchMock.defaultExpectation.results
		if results == nil {
			mmFetch.t.Fatal("No results are set for the JetFetcherMock.Fetch")
		}
		return (*results).ip1, (*results).err
	}
	if mmFetch.funcFetch != nil {
		return mmFetch.funcFetch(ctx, target, pulse)
	}
	mmFetch.t.Fatalf("Unexpected call to JetFetcherMock.Fetch. %v %v %v", ctx, target, pulse)
	return
}

// FetchAfterCounter returns a count of finished JetFetcherMock.Fetch invocations
func (mmFetch *JetFetcherMock) FetchAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFetch.afterFetchCounter)
}

// FetchBeforeCounter returns a count of JetFetcherMock.Fetch invocations
func (mmFetch *JetFetcherMock) FetchBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFetch.beforeFetchCounter)
}

// Calls returns a list of arguments used in each call to JetFetcherMock.Fetch.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFetch *mJetFetcherMockFetch) Calls() []*JetFetcherMockFetchParams {
	mmFetch.mutex.RLock()

	argCopy := make([]*JetFetcherMockFetchParams, len(mmFetch.callArgs))
	copy(argCopy, mmFetch.callArgs)

	mmFetch.mutex.RUnlock()

	return argCopy
}

// MinimockFetchDone returns true if the count of the Fetch invocations corresponds
// the number of defined expectations
func (m *JetFetcherMock) MinimockFetchDone() bool {
	for _, e := range m.FetchMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FetchMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFetchCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFetch != nil && mm_atomic.LoadUint64(&m.afterFetchCounter) < 1 {
		return false
	}
	return true
}

// MinimockFetchInspect logs each unmet expectation
func (m *JetFetcherMock) MinimockFetchInspect() {
	for _, e := range m.FetchMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetFetcherMock.Fetch with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FetchMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFetchCounter) < 1 {
		if m.FetchMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetFetcherMock.Fetch")
		} else {
			m.t.Errorf("Expected call to JetFetcherMock.Fetch with params: %#v", *m.FetchMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFetch != nil && mm_atomic.LoadUint64(&m.afterFetchCounter) < 1 {
		m.t.Error("Expected call to JetFetcherMock.Fetch")
	}
}

type mJetFetcherMockRelease struct {
	mock               *JetFetcherMock
	defaultExpectation *JetFetcherMockReleaseExpectation
	expectations       []*JetFetcherMockReleaseExpectation

	callArgs []*JetFetcherMockReleaseParams
	mutex    sync.RWMutex
}

// JetFetcherMockReleaseExpectation specifies expectation struct of the JetFetcher.Release
type JetFetcherMockReleaseExpectation struct {
	mock   *JetFetcherMock
	params *JetFetcherMockReleaseParams

	Counter uint64
}

// JetFetcherMockReleaseParams contains parameters of the JetFetcher.Release
type JetFetcherMockReleaseParams struct {
	ctx   context.Context
	jetID insolar.JetID
	pulse insolar.PulseNumber
}

// Expect sets up expected params for JetFetcher.Release
func (mmRelease *mJetFetcherMockRelease) Expect(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber) *mJetFetcherMockRelease {
	if mmRelease.mock.funcRelease != nil {
		mmRelease.mock.t.Fatalf("JetFetcherMock.Release mock is already set by Set")
	}

	if mmRelease.defaultExpectation == nil {
		mmRelease.defaultExpectation = &JetFetcherMockReleaseExpectation{}
	}

	mmRelease.defaultExpectation.params = &JetFetcherMockReleaseParams{ctx, jetID, pulse}
	for _, e := range mmRelease.expectations {
		if minimock.Equal(e.params, mmRelease.defaultExpectation.params) {
			mmRelease.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRelease.defaultExpectation.params)
		}
	}

	return mmRelease
}

// Inspect accepts an inspector function that has same arguments as the JetFetcher.Release
func (mmRelease *mJetFetcherMockRelease) Inspect(f func(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber)) *mJetFetcherMockRelease {
	if mmRelease.mock.inspectFuncRelease != nil {
		mmRelease.mock.t.Fatalf("Inspect function is already set for JetFetcherMock.Release")
	}

	mmRelease.mock.inspectFuncRelease = f

	return mmRelease
}

// Return sets up results that will be returned by JetFetcher.Release
func (mmRelease *mJetFetcherMockRelease) Return() *JetFetcherMock {
	if mmRelease.mock.funcRelease != nil {
		mmRelease.mock.t.Fatalf("JetFetcherMock.Release mock is already set by Set")
	}

	if mmRelease.defaultExpectation == nil {
		mmRelease.defaultExpectation = &JetFetcherMockReleaseExpectation{mock: mmRelease.mock}
	}

	return mmRelease.mock
}

//Set uses given function f to mock the JetFetcher.Release method
func (mmRelease *mJetFetcherMockRelease) Set(f func(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber)) *JetFetcherMock {
	if mmRelease.defaultExpectation != nil {
		mmRelease.mock.t.Fatalf("Default expectation is already set for the JetFetcher.Release method")
	}

	if len(mmRelease.expectations) > 0 {
		mmRelease.mock.t.Fatalf("Some expectations are already set for the JetFetcher.Release method")
	}

	mmRelease.mock.funcRelease = f
	return mmRelease.mock
}

// Release implements executor.JetFetcher
func (mmRelease *JetFetcherMock) Release(ctx context.Context, jetID insolar.JetID, pulse insolar.PulseNumber) {
	mm_atomic.AddUint64(&mmRelease.beforeReleaseCounter, 1)
	defer mm_atomic.AddUint64(&mmRelease.afterReleaseCounter, 1)

	if mmRelease.inspectFuncRelease != nil {
		mmRelease.inspectFuncRelease(ctx, jetID, pulse)
	}

	params := &JetFetcherMockReleaseParams{ctx, jetID, pulse}

	// Record call args
	mmRelease.ReleaseMock.mutex.Lock()
	mmRelease.ReleaseMock.callArgs = append(mmRelease.ReleaseMock.callArgs, params)
	mmRelease.ReleaseMock.mutex.Unlock()

	for _, e := range mmRelease.ReleaseMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmRelease.ReleaseMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRelease.ReleaseMock.defaultExpectation.Counter, 1)
		want := mmRelease.ReleaseMock.defaultExpectation.params
		got := JetFetcherMockReleaseParams{ctx, jetID, pulse}
		if want != nil && !minimock.Equal(*want, got) {
			mmRelease.t.Errorf("JetFetcherMock.Release got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		return

	}
	if mmRelease.funcRelease != nil {
		mmRelease.funcRelease(ctx, jetID, pulse)
		return
	}
	mmRelease.t.Fatalf("Unexpected call to JetFetcherMock.Release. %v %v %v", ctx, jetID, pulse)

}

// ReleaseAfterCounter returns a count of finished JetFetcherMock.Release invocations
func (mmRelease *JetFetcherMock) ReleaseAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRelease.afterReleaseCounter)
}

// ReleaseBeforeCounter returns a count of JetFetcherMock.Release invocations
func (mmRelease *JetFetcherMock) ReleaseBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRelease.beforeReleaseCounter)
}

// Calls returns a list of arguments used in each call to JetFetcherMock.Release.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRelease *mJetFetcherMockRelease) Calls() []*JetFetcherMockReleaseParams {
	mmRelease.mutex.RLock()

	argCopy := make([]*JetFetcherMockReleaseParams, len(mmRelease.callArgs))
	copy(argCopy, mmRelease.callArgs)

	mmRelease.mutex.RUnlock()

	return argCopy
}

// MinimockReleaseDone returns true if the count of the Release invocations corresponds
// the number of defined expectations
func (m *JetFetcherMock) MinimockReleaseDone() bool {
	for _, e := range m.ReleaseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReleaseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReleaseCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRelease != nil && mm_atomic.LoadUint64(&m.afterReleaseCounter) < 1 {
		return false
	}
	return true
}

// MinimockReleaseInspect logs each unmet expectation
func (m *JetFetcherMock) MinimockReleaseInspect() {
	for _, e := range m.ReleaseMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to JetFetcherMock.Release with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReleaseMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReleaseCounter) < 1 {
		if m.ReleaseMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to JetFetcherMock.Release")
		} else {
			m.t.Errorf("Expected call to JetFetcherMock.Release with params: %#v", *m.ReleaseMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRelease != nil && mm_atomic.LoadUint64(&m.afterReleaseCounter) < 1 {
		m.t.Error("Expected call to JetFetcherMock.Release")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *JetFetcherMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFetchInspect()

		m.MinimockReleaseInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *JetFetcherMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *JetFetcherMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFetchDone() &&
		m.MinimockReleaseDone()
}
