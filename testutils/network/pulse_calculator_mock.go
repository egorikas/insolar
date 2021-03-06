package network

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// PulseCalculatorMock implements storage.PulseCalculator
type PulseCalculatorMock struct {
	t minimock.Tester

	funcBackwards          func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)
	inspectFuncBackwards   func(ctx context.Context, pn insolar.PulseNumber, steps int)
	afterBackwardsCounter  uint64
	beforeBackwardsCounter uint64
	BackwardsMock          mPulseCalculatorMockBackwards

	funcForwards          func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)
	inspectFuncForwards   func(ctx context.Context, pn insolar.PulseNumber, steps int)
	afterForwardsCounter  uint64
	beforeForwardsCounter uint64
	ForwardsMock          mPulseCalculatorMockForwards
}

// NewPulseCalculatorMock returns a mock for storage.PulseCalculator
func NewPulseCalculatorMock(t minimock.Tester) *PulseCalculatorMock {
	m := &PulseCalculatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BackwardsMock = mPulseCalculatorMockBackwards{mock: m}
	m.BackwardsMock.callArgs = []*PulseCalculatorMockBackwardsParams{}

	m.ForwardsMock = mPulseCalculatorMockForwards{mock: m}
	m.ForwardsMock.callArgs = []*PulseCalculatorMockForwardsParams{}

	return m
}

type mPulseCalculatorMockBackwards struct {
	mock               *PulseCalculatorMock
	defaultExpectation *PulseCalculatorMockBackwardsExpectation
	expectations       []*PulseCalculatorMockBackwardsExpectation

	callArgs []*PulseCalculatorMockBackwardsParams
	mutex    sync.RWMutex
}

// PulseCalculatorMockBackwardsExpectation specifies expectation struct of the PulseCalculator.Backwards
type PulseCalculatorMockBackwardsExpectation struct {
	mock    *PulseCalculatorMock
	params  *PulseCalculatorMockBackwardsParams
	results *PulseCalculatorMockBackwardsResults
	Counter uint64
}

// PulseCalculatorMockBackwardsParams contains parameters of the PulseCalculator.Backwards
type PulseCalculatorMockBackwardsParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	steps int
}

// PulseCalculatorMockBackwardsResults contains results of the PulseCalculator.Backwards
type PulseCalculatorMockBackwardsResults struct {
	p1  insolar.Pulse
	err error
}

// Expect sets up expected params for PulseCalculator.Backwards
func (mmBackwards *mPulseCalculatorMockBackwards) Expect(ctx context.Context, pn insolar.PulseNumber, steps int) *mPulseCalculatorMockBackwards {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("PulseCalculatorMock.Backwards mock is already set by Set")
	}

	if mmBackwards.defaultExpectation == nil {
		mmBackwards.defaultExpectation = &PulseCalculatorMockBackwardsExpectation{}
	}

	mmBackwards.defaultExpectation.params = &PulseCalculatorMockBackwardsParams{ctx, pn, steps}
	for _, e := range mmBackwards.expectations {
		if minimock.Equal(e.params, mmBackwards.defaultExpectation.params) {
			mmBackwards.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmBackwards.defaultExpectation.params)
		}
	}

	return mmBackwards
}

// Inspect accepts an inspector function that has same arguments as the PulseCalculator.Backwards
func (mmBackwards *mPulseCalculatorMockBackwards) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, steps int)) *mPulseCalculatorMockBackwards {
	if mmBackwards.mock.inspectFuncBackwards != nil {
		mmBackwards.mock.t.Fatalf("Inspect function is already set for PulseCalculatorMock.Backwards")
	}

	mmBackwards.mock.inspectFuncBackwards = f

	return mmBackwards
}

// Return sets up results that will be returned by PulseCalculator.Backwards
func (mmBackwards *mPulseCalculatorMockBackwards) Return(p1 insolar.Pulse, err error) *PulseCalculatorMock {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("PulseCalculatorMock.Backwards mock is already set by Set")
	}

	if mmBackwards.defaultExpectation == nil {
		mmBackwards.defaultExpectation = &PulseCalculatorMockBackwardsExpectation{mock: mmBackwards.mock}
	}
	mmBackwards.defaultExpectation.results = &PulseCalculatorMockBackwardsResults{p1, err}
	return mmBackwards.mock
}

//Set uses given function f to mock the PulseCalculator.Backwards method
func (mmBackwards *mPulseCalculatorMockBackwards) Set(f func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)) *PulseCalculatorMock {
	if mmBackwards.defaultExpectation != nil {
		mmBackwards.mock.t.Fatalf("Default expectation is already set for the PulseCalculator.Backwards method")
	}

	if len(mmBackwards.expectations) > 0 {
		mmBackwards.mock.t.Fatalf("Some expectations are already set for the PulseCalculator.Backwards method")
	}

	mmBackwards.mock.funcBackwards = f
	return mmBackwards.mock
}

// When sets expectation for the PulseCalculator.Backwards which will trigger the result defined by the following
// Then helper
func (mmBackwards *mPulseCalculatorMockBackwards) When(ctx context.Context, pn insolar.PulseNumber, steps int) *PulseCalculatorMockBackwardsExpectation {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("PulseCalculatorMock.Backwards mock is already set by Set")
	}

	expectation := &PulseCalculatorMockBackwardsExpectation{
		mock:   mmBackwards.mock,
		params: &PulseCalculatorMockBackwardsParams{ctx, pn, steps},
	}
	mmBackwards.expectations = append(mmBackwards.expectations, expectation)
	return expectation
}

// Then sets up PulseCalculator.Backwards return parameters for the expectation previously defined by the When method
func (e *PulseCalculatorMockBackwardsExpectation) Then(p1 insolar.Pulse, err error) *PulseCalculatorMock {
	e.results = &PulseCalculatorMockBackwardsResults{p1, err}
	return e.mock
}

// Backwards implements storage.PulseCalculator
func (mmBackwards *PulseCalculatorMock) Backwards(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error) {
	mm_atomic.AddUint64(&mmBackwards.beforeBackwardsCounter, 1)
	defer mm_atomic.AddUint64(&mmBackwards.afterBackwardsCounter, 1)

	if mmBackwards.inspectFuncBackwards != nil {
		mmBackwards.inspectFuncBackwards(ctx, pn, steps)
	}

	params := &PulseCalculatorMockBackwardsParams{ctx, pn, steps}

	// Record call args
	mmBackwards.BackwardsMock.mutex.Lock()
	mmBackwards.BackwardsMock.callArgs = append(mmBackwards.BackwardsMock.callArgs, params)
	mmBackwards.BackwardsMock.mutex.Unlock()

	for _, e := range mmBackwards.BackwardsMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmBackwards.BackwardsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmBackwards.BackwardsMock.defaultExpectation.Counter, 1)
		want := mmBackwards.BackwardsMock.defaultExpectation.params
		got := PulseCalculatorMockBackwardsParams{ctx, pn, steps}
		if want != nil && !minimock.Equal(*want, got) {
			mmBackwards.t.Errorf("PulseCalculatorMock.Backwards got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmBackwards.BackwardsMock.defaultExpectation.results
		if results == nil {
			mmBackwards.t.Fatal("No results are set for the PulseCalculatorMock.Backwards")
		}
		return (*results).p1, (*results).err
	}
	if mmBackwards.funcBackwards != nil {
		return mmBackwards.funcBackwards(ctx, pn, steps)
	}
	mmBackwards.t.Fatalf("Unexpected call to PulseCalculatorMock.Backwards. %v %v %v", ctx, pn, steps)
	return
}

// BackwardsAfterCounter returns a count of finished PulseCalculatorMock.Backwards invocations
func (mmBackwards *PulseCalculatorMock) BackwardsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBackwards.afterBackwardsCounter)
}

// BackwardsBeforeCounter returns a count of PulseCalculatorMock.Backwards invocations
func (mmBackwards *PulseCalculatorMock) BackwardsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBackwards.beforeBackwardsCounter)
}

// Calls returns a list of arguments used in each call to PulseCalculatorMock.Backwards.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmBackwards *mPulseCalculatorMockBackwards) Calls() []*PulseCalculatorMockBackwardsParams {
	mmBackwards.mutex.RLock()

	argCopy := make([]*PulseCalculatorMockBackwardsParams, len(mmBackwards.callArgs))
	copy(argCopy, mmBackwards.callArgs)

	mmBackwards.mutex.RUnlock()

	return argCopy
}

// MinimockBackwardsDone returns true if the count of the Backwards invocations corresponds
// the number of defined expectations
func (m *PulseCalculatorMock) MinimockBackwardsDone() bool {
	for _, e := range m.BackwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BackwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBackwards != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		return false
	}
	return true
}

// MinimockBackwardsInspect logs each unmet expectation
func (m *PulseCalculatorMock) MinimockBackwardsInspect() {
	for _, e := range m.BackwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PulseCalculatorMock.Backwards with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BackwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		if m.BackwardsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PulseCalculatorMock.Backwards")
		} else {
			m.t.Errorf("Expected call to PulseCalculatorMock.Backwards with params: %#v", *m.BackwardsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBackwards != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		m.t.Error("Expected call to PulseCalculatorMock.Backwards")
	}
}

type mPulseCalculatorMockForwards struct {
	mock               *PulseCalculatorMock
	defaultExpectation *PulseCalculatorMockForwardsExpectation
	expectations       []*PulseCalculatorMockForwardsExpectation

	callArgs []*PulseCalculatorMockForwardsParams
	mutex    sync.RWMutex
}

// PulseCalculatorMockForwardsExpectation specifies expectation struct of the PulseCalculator.Forwards
type PulseCalculatorMockForwardsExpectation struct {
	mock    *PulseCalculatorMock
	params  *PulseCalculatorMockForwardsParams
	results *PulseCalculatorMockForwardsResults
	Counter uint64
}

// PulseCalculatorMockForwardsParams contains parameters of the PulseCalculator.Forwards
type PulseCalculatorMockForwardsParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	steps int
}

// PulseCalculatorMockForwardsResults contains results of the PulseCalculator.Forwards
type PulseCalculatorMockForwardsResults struct {
	p1  insolar.Pulse
	err error
}

// Expect sets up expected params for PulseCalculator.Forwards
func (mmForwards *mPulseCalculatorMockForwards) Expect(ctx context.Context, pn insolar.PulseNumber, steps int) *mPulseCalculatorMockForwards {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("PulseCalculatorMock.Forwards mock is already set by Set")
	}

	if mmForwards.defaultExpectation == nil {
		mmForwards.defaultExpectation = &PulseCalculatorMockForwardsExpectation{}
	}

	mmForwards.defaultExpectation.params = &PulseCalculatorMockForwardsParams{ctx, pn, steps}
	for _, e := range mmForwards.expectations {
		if minimock.Equal(e.params, mmForwards.defaultExpectation.params) {
			mmForwards.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForwards.defaultExpectation.params)
		}
	}

	return mmForwards
}

// Inspect accepts an inspector function that has same arguments as the PulseCalculator.Forwards
func (mmForwards *mPulseCalculatorMockForwards) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, steps int)) *mPulseCalculatorMockForwards {
	if mmForwards.mock.inspectFuncForwards != nil {
		mmForwards.mock.t.Fatalf("Inspect function is already set for PulseCalculatorMock.Forwards")
	}

	mmForwards.mock.inspectFuncForwards = f

	return mmForwards
}

// Return sets up results that will be returned by PulseCalculator.Forwards
func (mmForwards *mPulseCalculatorMockForwards) Return(p1 insolar.Pulse, err error) *PulseCalculatorMock {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("PulseCalculatorMock.Forwards mock is already set by Set")
	}

	if mmForwards.defaultExpectation == nil {
		mmForwards.defaultExpectation = &PulseCalculatorMockForwardsExpectation{mock: mmForwards.mock}
	}
	mmForwards.defaultExpectation.results = &PulseCalculatorMockForwardsResults{p1, err}
	return mmForwards.mock
}

//Set uses given function f to mock the PulseCalculator.Forwards method
func (mmForwards *mPulseCalculatorMockForwards) Set(f func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)) *PulseCalculatorMock {
	if mmForwards.defaultExpectation != nil {
		mmForwards.mock.t.Fatalf("Default expectation is already set for the PulseCalculator.Forwards method")
	}

	if len(mmForwards.expectations) > 0 {
		mmForwards.mock.t.Fatalf("Some expectations are already set for the PulseCalculator.Forwards method")
	}

	mmForwards.mock.funcForwards = f
	return mmForwards.mock
}

// When sets expectation for the PulseCalculator.Forwards which will trigger the result defined by the following
// Then helper
func (mmForwards *mPulseCalculatorMockForwards) When(ctx context.Context, pn insolar.PulseNumber, steps int) *PulseCalculatorMockForwardsExpectation {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("PulseCalculatorMock.Forwards mock is already set by Set")
	}

	expectation := &PulseCalculatorMockForwardsExpectation{
		mock:   mmForwards.mock,
		params: &PulseCalculatorMockForwardsParams{ctx, pn, steps},
	}
	mmForwards.expectations = append(mmForwards.expectations, expectation)
	return expectation
}

// Then sets up PulseCalculator.Forwards return parameters for the expectation previously defined by the When method
func (e *PulseCalculatorMockForwardsExpectation) Then(p1 insolar.Pulse, err error) *PulseCalculatorMock {
	e.results = &PulseCalculatorMockForwardsResults{p1, err}
	return e.mock
}

// Forwards implements storage.PulseCalculator
func (mmForwards *PulseCalculatorMock) Forwards(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error) {
	mm_atomic.AddUint64(&mmForwards.beforeForwardsCounter, 1)
	defer mm_atomic.AddUint64(&mmForwards.afterForwardsCounter, 1)

	if mmForwards.inspectFuncForwards != nil {
		mmForwards.inspectFuncForwards(ctx, pn, steps)
	}

	params := &PulseCalculatorMockForwardsParams{ctx, pn, steps}

	// Record call args
	mmForwards.ForwardsMock.mutex.Lock()
	mmForwards.ForwardsMock.callArgs = append(mmForwards.ForwardsMock.callArgs, params)
	mmForwards.ForwardsMock.mutex.Unlock()

	for _, e := range mmForwards.ForwardsMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.p1, e.results.err
		}
	}

	if mmForwards.ForwardsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmForwards.ForwardsMock.defaultExpectation.Counter, 1)
		want := mmForwards.ForwardsMock.defaultExpectation.params
		got := PulseCalculatorMockForwardsParams{ctx, pn, steps}
		if want != nil && !minimock.Equal(*want, got) {
			mmForwards.t.Errorf("PulseCalculatorMock.Forwards got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForwards.ForwardsMock.defaultExpectation.results
		if results == nil {
			mmForwards.t.Fatal("No results are set for the PulseCalculatorMock.Forwards")
		}
		return (*results).p1, (*results).err
	}
	if mmForwards.funcForwards != nil {
		return mmForwards.funcForwards(ctx, pn, steps)
	}
	mmForwards.t.Fatalf("Unexpected call to PulseCalculatorMock.Forwards. %v %v %v", ctx, pn, steps)
	return
}

// ForwardsAfterCounter returns a count of finished PulseCalculatorMock.Forwards invocations
func (mmForwards *PulseCalculatorMock) ForwardsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForwards.afterForwardsCounter)
}

// ForwardsBeforeCounter returns a count of PulseCalculatorMock.Forwards invocations
func (mmForwards *PulseCalculatorMock) ForwardsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForwards.beforeForwardsCounter)
}

// Calls returns a list of arguments used in each call to PulseCalculatorMock.Forwards.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForwards *mPulseCalculatorMockForwards) Calls() []*PulseCalculatorMockForwardsParams {
	mmForwards.mutex.RLock()

	argCopy := make([]*PulseCalculatorMockForwardsParams, len(mmForwards.callArgs))
	copy(argCopy, mmForwards.callArgs)

	mmForwards.mutex.RUnlock()

	return argCopy
}

// MinimockForwardsDone returns true if the count of the Forwards invocations corresponds
// the number of defined expectations
func (m *PulseCalculatorMock) MinimockForwardsDone() bool {
	for _, e := range m.ForwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForwards != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		return false
	}
	return true
}

// MinimockForwardsInspect logs each unmet expectation
func (m *PulseCalculatorMock) MinimockForwardsInspect() {
	for _, e := range m.ForwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to PulseCalculatorMock.Forwards with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		if m.ForwardsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to PulseCalculatorMock.Forwards")
		} else {
			m.t.Errorf("Expected call to PulseCalculatorMock.Forwards with params: %#v", *m.ForwardsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForwards != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		m.t.Error("Expected call to PulseCalculatorMock.Forwards")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *PulseCalculatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockBackwardsInspect()

		m.MinimockForwardsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *PulseCalculatorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *PulseCalculatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockBackwardsDone() &&
		m.MinimockForwardsDone()
}
