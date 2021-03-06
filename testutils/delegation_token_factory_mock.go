package testutils

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	mm_insolar "github.com/insolar/insolar/insolar"
)

// DelegationTokenFactoryMock implements insolar.DelegationTokenFactory
type DelegationTokenFactoryMock struct {
	t minimock.Tester

	funcIssuePendingExecution          func(msg mm_insolar.Message, pulse mm_insolar.PulseNumber) (d1 mm_insolar.DelegationToken, err error)
	inspectFuncIssuePendingExecution   func(msg mm_insolar.Message, pulse mm_insolar.PulseNumber)
	afterIssuePendingExecutionCounter  uint64
	beforeIssuePendingExecutionCounter uint64
	IssuePendingExecutionMock          mDelegationTokenFactoryMockIssuePendingExecution

	funcVerify          func(parcel mm_insolar.Parcel) (b1 bool, err error)
	inspectFuncVerify   func(parcel mm_insolar.Parcel)
	afterVerifyCounter  uint64
	beforeVerifyCounter uint64
	VerifyMock          mDelegationTokenFactoryMockVerify
}

// NewDelegationTokenFactoryMock returns a mock for insolar.DelegationTokenFactory
func NewDelegationTokenFactoryMock(t minimock.Tester) *DelegationTokenFactoryMock {
	m := &DelegationTokenFactoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.IssuePendingExecutionMock = mDelegationTokenFactoryMockIssuePendingExecution{mock: m}
	m.IssuePendingExecutionMock.callArgs = []*DelegationTokenFactoryMockIssuePendingExecutionParams{}

	m.VerifyMock = mDelegationTokenFactoryMockVerify{mock: m}
	m.VerifyMock.callArgs = []*DelegationTokenFactoryMockVerifyParams{}

	return m
}

type mDelegationTokenFactoryMockIssuePendingExecution struct {
	mock               *DelegationTokenFactoryMock
	defaultExpectation *DelegationTokenFactoryMockIssuePendingExecutionExpectation
	expectations       []*DelegationTokenFactoryMockIssuePendingExecutionExpectation

	callArgs []*DelegationTokenFactoryMockIssuePendingExecutionParams
	mutex    sync.RWMutex
}

// DelegationTokenFactoryMockIssuePendingExecutionExpectation specifies expectation struct of the DelegationTokenFactory.IssuePendingExecution
type DelegationTokenFactoryMockIssuePendingExecutionExpectation struct {
	mock    *DelegationTokenFactoryMock
	params  *DelegationTokenFactoryMockIssuePendingExecutionParams
	results *DelegationTokenFactoryMockIssuePendingExecutionResults
	Counter uint64
}

// DelegationTokenFactoryMockIssuePendingExecutionParams contains parameters of the DelegationTokenFactory.IssuePendingExecution
type DelegationTokenFactoryMockIssuePendingExecutionParams struct {
	msg   mm_insolar.Message
	pulse mm_insolar.PulseNumber
}

// DelegationTokenFactoryMockIssuePendingExecutionResults contains results of the DelegationTokenFactory.IssuePendingExecution
type DelegationTokenFactoryMockIssuePendingExecutionResults struct {
	d1  mm_insolar.DelegationToken
	err error
}

// Expect sets up expected params for DelegationTokenFactory.IssuePendingExecution
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) Expect(msg mm_insolar.Message, pulse mm_insolar.PulseNumber) *mDelegationTokenFactoryMockIssuePendingExecution {
	if mmIssuePendingExecution.mock.funcIssuePendingExecution != nil {
		mmIssuePendingExecution.mock.t.Fatalf("DelegationTokenFactoryMock.IssuePendingExecution mock is already set by Set")
	}

	if mmIssuePendingExecution.defaultExpectation == nil {
		mmIssuePendingExecution.defaultExpectation = &DelegationTokenFactoryMockIssuePendingExecutionExpectation{}
	}

	mmIssuePendingExecution.defaultExpectation.params = &DelegationTokenFactoryMockIssuePendingExecutionParams{msg, pulse}
	for _, e := range mmIssuePendingExecution.expectations {
		if minimock.Equal(e.params, mmIssuePendingExecution.defaultExpectation.params) {
			mmIssuePendingExecution.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIssuePendingExecution.defaultExpectation.params)
		}
	}

	return mmIssuePendingExecution
}

// Inspect accepts an inspector function that has same arguments as the DelegationTokenFactory.IssuePendingExecution
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) Inspect(f func(msg mm_insolar.Message, pulse mm_insolar.PulseNumber)) *mDelegationTokenFactoryMockIssuePendingExecution {
	if mmIssuePendingExecution.mock.inspectFuncIssuePendingExecution != nil {
		mmIssuePendingExecution.mock.t.Fatalf("Inspect function is already set for DelegationTokenFactoryMock.IssuePendingExecution")
	}

	mmIssuePendingExecution.mock.inspectFuncIssuePendingExecution = f

	return mmIssuePendingExecution
}

// Return sets up results that will be returned by DelegationTokenFactory.IssuePendingExecution
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) Return(d1 mm_insolar.DelegationToken, err error) *DelegationTokenFactoryMock {
	if mmIssuePendingExecution.mock.funcIssuePendingExecution != nil {
		mmIssuePendingExecution.mock.t.Fatalf("DelegationTokenFactoryMock.IssuePendingExecution mock is already set by Set")
	}

	if mmIssuePendingExecution.defaultExpectation == nil {
		mmIssuePendingExecution.defaultExpectation = &DelegationTokenFactoryMockIssuePendingExecutionExpectation{mock: mmIssuePendingExecution.mock}
	}
	mmIssuePendingExecution.defaultExpectation.results = &DelegationTokenFactoryMockIssuePendingExecutionResults{d1, err}
	return mmIssuePendingExecution.mock
}

//Set uses given function f to mock the DelegationTokenFactory.IssuePendingExecution method
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) Set(f func(msg mm_insolar.Message, pulse mm_insolar.PulseNumber) (d1 mm_insolar.DelegationToken, err error)) *DelegationTokenFactoryMock {
	if mmIssuePendingExecution.defaultExpectation != nil {
		mmIssuePendingExecution.mock.t.Fatalf("Default expectation is already set for the DelegationTokenFactory.IssuePendingExecution method")
	}

	if len(mmIssuePendingExecution.expectations) > 0 {
		mmIssuePendingExecution.mock.t.Fatalf("Some expectations are already set for the DelegationTokenFactory.IssuePendingExecution method")
	}

	mmIssuePendingExecution.mock.funcIssuePendingExecution = f
	return mmIssuePendingExecution.mock
}

// When sets expectation for the DelegationTokenFactory.IssuePendingExecution which will trigger the result defined by the following
// Then helper
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) When(msg mm_insolar.Message, pulse mm_insolar.PulseNumber) *DelegationTokenFactoryMockIssuePendingExecutionExpectation {
	if mmIssuePendingExecution.mock.funcIssuePendingExecution != nil {
		mmIssuePendingExecution.mock.t.Fatalf("DelegationTokenFactoryMock.IssuePendingExecution mock is already set by Set")
	}

	expectation := &DelegationTokenFactoryMockIssuePendingExecutionExpectation{
		mock:   mmIssuePendingExecution.mock,
		params: &DelegationTokenFactoryMockIssuePendingExecutionParams{msg, pulse},
	}
	mmIssuePendingExecution.expectations = append(mmIssuePendingExecution.expectations, expectation)
	return expectation
}

// Then sets up DelegationTokenFactory.IssuePendingExecution return parameters for the expectation previously defined by the When method
func (e *DelegationTokenFactoryMockIssuePendingExecutionExpectation) Then(d1 mm_insolar.DelegationToken, err error) *DelegationTokenFactoryMock {
	e.results = &DelegationTokenFactoryMockIssuePendingExecutionResults{d1, err}
	return e.mock
}

// IssuePendingExecution implements insolar.DelegationTokenFactory
func (mmIssuePendingExecution *DelegationTokenFactoryMock) IssuePendingExecution(msg mm_insolar.Message, pulse mm_insolar.PulseNumber) (d1 mm_insolar.DelegationToken, err error) {
	mm_atomic.AddUint64(&mmIssuePendingExecution.beforeIssuePendingExecutionCounter, 1)
	defer mm_atomic.AddUint64(&mmIssuePendingExecution.afterIssuePendingExecutionCounter, 1)

	if mmIssuePendingExecution.inspectFuncIssuePendingExecution != nil {
		mmIssuePendingExecution.inspectFuncIssuePendingExecution(msg, pulse)
	}

	params := &DelegationTokenFactoryMockIssuePendingExecutionParams{msg, pulse}

	// Record call args
	mmIssuePendingExecution.IssuePendingExecutionMock.mutex.Lock()
	mmIssuePendingExecution.IssuePendingExecutionMock.callArgs = append(mmIssuePendingExecution.IssuePendingExecutionMock.callArgs, params)
	mmIssuePendingExecution.IssuePendingExecutionMock.mutex.Unlock()

	for _, e := range mmIssuePendingExecution.IssuePendingExecutionMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.d1, e.results.err
		}
	}

	if mmIssuePendingExecution.IssuePendingExecutionMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIssuePendingExecution.IssuePendingExecutionMock.defaultExpectation.Counter, 1)
		want := mmIssuePendingExecution.IssuePendingExecutionMock.defaultExpectation.params
		got := DelegationTokenFactoryMockIssuePendingExecutionParams{msg, pulse}
		if want != nil && !minimock.Equal(*want, got) {
			mmIssuePendingExecution.t.Errorf("DelegationTokenFactoryMock.IssuePendingExecution got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmIssuePendingExecution.IssuePendingExecutionMock.defaultExpectation.results
		if results == nil {
			mmIssuePendingExecution.t.Fatal("No results are set for the DelegationTokenFactoryMock.IssuePendingExecution")
		}
		return (*results).d1, (*results).err
	}
	if mmIssuePendingExecution.funcIssuePendingExecution != nil {
		return mmIssuePendingExecution.funcIssuePendingExecution(msg, pulse)
	}
	mmIssuePendingExecution.t.Fatalf("Unexpected call to DelegationTokenFactoryMock.IssuePendingExecution. %v %v", msg, pulse)
	return
}

// IssuePendingExecutionAfterCounter returns a count of finished DelegationTokenFactoryMock.IssuePendingExecution invocations
func (mmIssuePendingExecution *DelegationTokenFactoryMock) IssuePendingExecutionAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIssuePendingExecution.afterIssuePendingExecutionCounter)
}

// IssuePendingExecutionBeforeCounter returns a count of DelegationTokenFactoryMock.IssuePendingExecution invocations
func (mmIssuePendingExecution *DelegationTokenFactoryMock) IssuePendingExecutionBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIssuePendingExecution.beforeIssuePendingExecutionCounter)
}

// Calls returns a list of arguments used in each call to DelegationTokenFactoryMock.IssuePendingExecution.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIssuePendingExecution *mDelegationTokenFactoryMockIssuePendingExecution) Calls() []*DelegationTokenFactoryMockIssuePendingExecutionParams {
	mmIssuePendingExecution.mutex.RLock()

	argCopy := make([]*DelegationTokenFactoryMockIssuePendingExecutionParams, len(mmIssuePendingExecution.callArgs))
	copy(argCopy, mmIssuePendingExecution.callArgs)

	mmIssuePendingExecution.mutex.RUnlock()

	return argCopy
}

// MinimockIssuePendingExecutionDone returns true if the count of the IssuePendingExecution invocations corresponds
// the number of defined expectations
func (m *DelegationTokenFactoryMock) MinimockIssuePendingExecutionDone() bool {
	for _, e := range m.IssuePendingExecutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IssuePendingExecutionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIssuePendingExecutionCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIssuePendingExecution != nil && mm_atomic.LoadUint64(&m.afterIssuePendingExecutionCounter) < 1 {
		return false
	}
	return true
}

// MinimockIssuePendingExecutionInspect logs each unmet expectation
func (m *DelegationTokenFactoryMock) MinimockIssuePendingExecutionInspect() {
	for _, e := range m.IssuePendingExecutionMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DelegationTokenFactoryMock.IssuePendingExecution with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IssuePendingExecutionMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIssuePendingExecutionCounter) < 1 {
		if m.IssuePendingExecutionMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DelegationTokenFactoryMock.IssuePendingExecution")
		} else {
			m.t.Errorf("Expected call to DelegationTokenFactoryMock.IssuePendingExecution with params: %#v", *m.IssuePendingExecutionMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIssuePendingExecution != nil && mm_atomic.LoadUint64(&m.afterIssuePendingExecutionCounter) < 1 {
		m.t.Error("Expected call to DelegationTokenFactoryMock.IssuePendingExecution")
	}
}

type mDelegationTokenFactoryMockVerify struct {
	mock               *DelegationTokenFactoryMock
	defaultExpectation *DelegationTokenFactoryMockVerifyExpectation
	expectations       []*DelegationTokenFactoryMockVerifyExpectation

	callArgs []*DelegationTokenFactoryMockVerifyParams
	mutex    sync.RWMutex
}

// DelegationTokenFactoryMockVerifyExpectation specifies expectation struct of the DelegationTokenFactory.Verify
type DelegationTokenFactoryMockVerifyExpectation struct {
	mock    *DelegationTokenFactoryMock
	params  *DelegationTokenFactoryMockVerifyParams
	results *DelegationTokenFactoryMockVerifyResults
	Counter uint64
}

// DelegationTokenFactoryMockVerifyParams contains parameters of the DelegationTokenFactory.Verify
type DelegationTokenFactoryMockVerifyParams struct {
	parcel mm_insolar.Parcel
}

// DelegationTokenFactoryMockVerifyResults contains results of the DelegationTokenFactory.Verify
type DelegationTokenFactoryMockVerifyResults struct {
	b1  bool
	err error
}

// Expect sets up expected params for DelegationTokenFactory.Verify
func (mmVerify *mDelegationTokenFactoryMockVerify) Expect(parcel mm_insolar.Parcel) *mDelegationTokenFactoryMockVerify {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("DelegationTokenFactoryMock.Verify mock is already set by Set")
	}

	if mmVerify.defaultExpectation == nil {
		mmVerify.defaultExpectation = &DelegationTokenFactoryMockVerifyExpectation{}
	}

	mmVerify.defaultExpectation.params = &DelegationTokenFactoryMockVerifyParams{parcel}
	for _, e := range mmVerify.expectations {
		if minimock.Equal(e.params, mmVerify.defaultExpectation.params) {
			mmVerify.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmVerify.defaultExpectation.params)
		}
	}

	return mmVerify
}

// Inspect accepts an inspector function that has same arguments as the DelegationTokenFactory.Verify
func (mmVerify *mDelegationTokenFactoryMockVerify) Inspect(f func(parcel mm_insolar.Parcel)) *mDelegationTokenFactoryMockVerify {
	if mmVerify.mock.inspectFuncVerify != nil {
		mmVerify.mock.t.Fatalf("Inspect function is already set for DelegationTokenFactoryMock.Verify")
	}

	mmVerify.mock.inspectFuncVerify = f

	return mmVerify
}

// Return sets up results that will be returned by DelegationTokenFactory.Verify
func (mmVerify *mDelegationTokenFactoryMockVerify) Return(b1 bool, err error) *DelegationTokenFactoryMock {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("DelegationTokenFactoryMock.Verify mock is already set by Set")
	}

	if mmVerify.defaultExpectation == nil {
		mmVerify.defaultExpectation = &DelegationTokenFactoryMockVerifyExpectation{mock: mmVerify.mock}
	}
	mmVerify.defaultExpectation.results = &DelegationTokenFactoryMockVerifyResults{b1, err}
	return mmVerify.mock
}

//Set uses given function f to mock the DelegationTokenFactory.Verify method
func (mmVerify *mDelegationTokenFactoryMockVerify) Set(f func(parcel mm_insolar.Parcel) (b1 bool, err error)) *DelegationTokenFactoryMock {
	if mmVerify.defaultExpectation != nil {
		mmVerify.mock.t.Fatalf("Default expectation is already set for the DelegationTokenFactory.Verify method")
	}

	if len(mmVerify.expectations) > 0 {
		mmVerify.mock.t.Fatalf("Some expectations are already set for the DelegationTokenFactory.Verify method")
	}

	mmVerify.mock.funcVerify = f
	return mmVerify.mock
}

// When sets expectation for the DelegationTokenFactory.Verify which will trigger the result defined by the following
// Then helper
func (mmVerify *mDelegationTokenFactoryMockVerify) When(parcel mm_insolar.Parcel) *DelegationTokenFactoryMockVerifyExpectation {
	if mmVerify.mock.funcVerify != nil {
		mmVerify.mock.t.Fatalf("DelegationTokenFactoryMock.Verify mock is already set by Set")
	}

	expectation := &DelegationTokenFactoryMockVerifyExpectation{
		mock:   mmVerify.mock,
		params: &DelegationTokenFactoryMockVerifyParams{parcel},
	}
	mmVerify.expectations = append(mmVerify.expectations, expectation)
	return expectation
}

// Then sets up DelegationTokenFactory.Verify return parameters for the expectation previously defined by the When method
func (e *DelegationTokenFactoryMockVerifyExpectation) Then(b1 bool, err error) *DelegationTokenFactoryMock {
	e.results = &DelegationTokenFactoryMockVerifyResults{b1, err}
	return e.mock
}

// Verify implements insolar.DelegationTokenFactory
func (mmVerify *DelegationTokenFactoryMock) Verify(parcel mm_insolar.Parcel) (b1 bool, err error) {
	mm_atomic.AddUint64(&mmVerify.beforeVerifyCounter, 1)
	defer mm_atomic.AddUint64(&mmVerify.afterVerifyCounter, 1)

	if mmVerify.inspectFuncVerify != nil {
		mmVerify.inspectFuncVerify(parcel)
	}

	params := &DelegationTokenFactoryMockVerifyParams{parcel}

	// Record call args
	mmVerify.VerifyMock.mutex.Lock()
	mmVerify.VerifyMock.callArgs = append(mmVerify.VerifyMock.callArgs, params)
	mmVerify.VerifyMock.mutex.Unlock()

	for _, e := range mmVerify.VerifyMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.b1, e.results.err
		}
	}

	if mmVerify.VerifyMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmVerify.VerifyMock.defaultExpectation.Counter, 1)
		want := mmVerify.VerifyMock.defaultExpectation.params
		got := DelegationTokenFactoryMockVerifyParams{parcel}
		if want != nil && !minimock.Equal(*want, got) {
			mmVerify.t.Errorf("DelegationTokenFactoryMock.Verify got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmVerify.VerifyMock.defaultExpectation.results
		if results == nil {
			mmVerify.t.Fatal("No results are set for the DelegationTokenFactoryMock.Verify")
		}
		return (*results).b1, (*results).err
	}
	if mmVerify.funcVerify != nil {
		return mmVerify.funcVerify(parcel)
	}
	mmVerify.t.Fatalf("Unexpected call to DelegationTokenFactoryMock.Verify. %v", parcel)
	return
}

// VerifyAfterCounter returns a count of finished DelegationTokenFactoryMock.Verify invocations
func (mmVerify *DelegationTokenFactoryMock) VerifyAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVerify.afterVerifyCounter)
}

// VerifyBeforeCounter returns a count of DelegationTokenFactoryMock.Verify invocations
func (mmVerify *DelegationTokenFactoryMock) VerifyBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmVerify.beforeVerifyCounter)
}

// Calls returns a list of arguments used in each call to DelegationTokenFactoryMock.Verify.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmVerify *mDelegationTokenFactoryMockVerify) Calls() []*DelegationTokenFactoryMockVerifyParams {
	mmVerify.mutex.RLock()

	argCopy := make([]*DelegationTokenFactoryMockVerifyParams, len(mmVerify.callArgs))
	copy(argCopy, mmVerify.callArgs)

	mmVerify.mutex.RUnlock()

	return argCopy
}

// MinimockVerifyDone returns true if the count of the Verify invocations corresponds
// the number of defined expectations
func (m *DelegationTokenFactoryMock) MinimockVerifyDone() bool {
	for _, e := range m.VerifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VerifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVerify != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		return false
	}
	return true
}

// MinimockVerifyInspect logs each unmet expectation
func (m *DelegationTokenFactoryMock) MinimockVerifyInspect() {
	for _, e := range m.VerifyMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DelegationTokenFactoryMock.Verify with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.VerifyMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		if m.VerifyMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DelegationTokenFactoryMock.Verify")
		} else {
			m.t.Errorf("Expected call to DelegationTokenFactoryMock.Verify with params: %#v", *m.VerifyMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcVerify != nil && mm_atomic.LoadUint64(&m.afterVerifyCounter) < 1 {
		m.t.Error("Expected call to DelegationTokenFactoryMock.Verify")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DelegationTokenFactoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockIssuePendingExecutionInspect()

		m.MinimockVerifyInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DelegationTokenFactoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *DelegationTokenFactoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockIssuePendingExecutionDone() &&
		m.MinimockVerifyDone()
}
