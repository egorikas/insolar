package pulse

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
)

// CalculatorMock implements pulse.Calculator
type CalculatorMock struct {
	t minimock.Tester

	funcBackwards          func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)
	inspectFuncBackwards   func(ctx context.Context, pn insolar.PulseNumber, steps int)
	afterBackwardsCounter  uint64
	beforeBackwardsCounter uint64
	BackwardsMock          mCalculatorMockBackwards

	funcForwards          func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)
	inspectFuncForwards   func(ctx context.Context, pn insolar.PulseNumber, steps int)
	afterForwardsCounter  uint64
	beforeForwardsCounter uint64
	ForwardsMock          mCalculatorMockForwards
}

// NewCalculatorMock returns a mock for pulse.Calculator
func NewCalculatorMock(t minimock.Tester) *CalculatorMock {
	m := &CalculatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BackwardsMock = mCalculatorMockBackwards{mock: m}
	m.BackwardsMock.callArgs = []*CalculatorMockBackwardsParams{}

	m.ForwardsMock = mCalculatorMockForwards{mock: m}
	m.ForwardsMock.callArgs = []*CalculatorMockForwardsParams{}

	return m
}

type mCalculatorMockBackwards struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockBackwardsExpectation
	expectations       []*CalculatorMockBackwardsExpectation

	callArgs []*CalculatorMockBackwardsParams
	mutex    sync.RWMutex
}

// CalculatorMockBackwardsExpectation specifies expectation struct of the Calculator.Backwards
type CalculatorMockBackwardsExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockBackwardsParams
	results *CalculatorMockBackwardsResults
	Counter uint64
}

// CalculatorMockBackwardsParams contains parameters of the Calculator.Backwards
type CalculatorMockBackwardsParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	steps int
}

// CalculatorMockBackwardsResults contains results of the Calculator.Backwards
type CalculatorMockBackwardsResults struct {
	p1  insolar.Pulse
	err error
}

// Expect sets up expected params for Calculator.Backwards
func (mmBackwards *mCalculatorMockBackwards) Expect(ctx context.Context, pn insolar.PulseNumber, steps int) *mCalculatorMockBackwards {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("CalculatorMock.Backwards mock is already set by Set")
	}

	if mmBackwards.defaultExpectation == nil {
		mmBackwards.defaultExpectation = &CalculatorMockBackwardsExpectation{}
	}

	mmBackwards.defaultExpectation.params = &CalculatorMockBackwardsParams{ctx, pn, steps}
	for _, e := range mmBackwards.expectations {
		if minimock.Equal(e.params, mmBackwards.defaultExpectation.params) {
			mmBackwards.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmBackwards.defaultExpectation.params)
		}
	}

	return mmBackwards
}

// Inspect accepts an inspector function that has same arguments as the Calculator.Backwards
func (mmBackwards *mCalculatorMockBackwards) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, steps int)) *mCalculatorMockBackwards {
	if mmBackwards.mock.inspectFuncBackwards != nil {
		mmBackwards.mock.t.Fatalf("Inspect function is already set for CalculatorMock.Backwards")
	}

	mmBackwards.mock.inspectFuncBackwards = f

	return mmBackwards
}

// Return sets up results that will be returned by Calculator.Backwards
func (mmBackwards *mCalculatorMockBackwards) Return(p1 insolar.Pulse, err error) *CalculatorMock {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("CalculatorMock.Backwards mock is already set by Set")
	}

	if mmBackwards.defaultExpectation == nil {
		mmBackwards.defaultExpectation = &CalculatorMockBackwardsExpectation{mock: mmBackwards.mock}
	}
	mmBackwards.defaultExpectation.results = &CalculatorMockBackwardsResults{p1, err}
	return mmBackwards.mock
}

//Set uses given function f to mock the Calculator.Backwards method
func (mmBackwards *mCalculatorMockBackwards) Set(f func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)) *CalculatorMock {
	if mmBackwards.defaultExpectation != nil {
		mmBackwards.mock.t.Fatalf("Default expectation is already set for the Calculator.Backwards method")
	}

	if len(mmBackwards.expectations) > 0 {
		mmBackwards.mock.t.Fatalf("Some expectations are already set for the Calculator.Backwards method")
	}

	mmBackwards.mock.funcBackwards = f
	return mmBackwards.mock
}

// When sets expectation for the Calculator.Backwards which will trigger the result defined by the following
// Then helper
func (mmBackwards *mCalculatorMockBackwards) When(ctx context.Context, pn insolar.PulseNumber, steps int) *CalculatorMockBackwardsExpectation {
	if mmBackwards.mock.funcBackwards != nil {
		mmBackwards.mock.t.Fatalf("CalculatorMock.Backwards mock is already set by Set")
	}

	expectation := &CalculatorMockBackwardsExpectation{
		mock:   mmBackwards.mock,
		params: &CalculatorMockBackwardsParams{ctx, pn, steps},
	}
	mmBackwards.expectations = append(mmBackwards.expectations, expectation)
	return expectation
}

// Then sets up Calculator.Backwards return parameters for the expectation previously defined by the When method
func (e *CalculatorMockBackwardsExpectation) Then(p1 insolar.Pulse, err error) *CalculatorMock {
	e.results = &CalculatorMockBackwardsResults{p1, err}
	return e.mock
}

// Backwards implements pulse.Calculator
func (mmBackwards *CalculatorMock) Backwards(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error) {
	mm_atomic.AddUint64(&mmBackwards.beforeBackwardsCounter, 1)
	defer mm_atomic.AddUint64(&mmBackwards.afterBackwardsCounter, 1)

	if mmBackwards.inspectFuncBackwards != nil {
		mmBackwards.inspectFuncBackwards(ctx, pn, steps)
	}

	params := &CalculatorMockBackwardsParams{ctx, pn, steps}

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
		got := CalculatorMockBackwardsParams{ctx, pn, steps}
		if want != nil && !minimock.Equal(*want, got) {
			mmBackwards.t.Errorf("CalculatorMock.Backwards got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmBackwards.BackwardsMock.defaultExpectation.results
		if results == nil {
			mmBackwards.t.Fatal("No results are set for the CalculatorMock.Backwards")
		}
		return (*results).p1, (*results).err
	}
	if mmBackwards.funcBackwards != nil {
		return mmBackwards.funcBackwards(ctx, pn, steps)
	}
	mmBackwards.t.Fatalf("Unexpected call to CalculatorMock.Backwards. %v %v %v", ctx, pn, steps)
	return
}

// BackwardsAfterCounter returns a count of finished CalculatorMock.Backwards invocations
func (mmBackwards *CalculatorMock) BackwardsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBackwards.afterBackwardsCounter)
}

// BackwardsBeforeCounter returns a count of CalculatorMock.Backwards invocations
func (mmBackwards *CalculatorMock) BackwardsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBackwards.beforeBackwardsCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.Backwards.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmBackwards *mCalculatorMockBackwards) Calls() []*CalculatorMockBackwardsParams {
	mmBackwards.mutex.RLock()

	argCopy := make([]*CalculatorMockBackwardsParams, len(mmBackwards.callArgs))
	copy(argCopy, mmBackwards.callArgs)

	mmBackwards.mutex.RUnlock()

	return argCopy
}

// MinimockBackwardsDone returns true if the count of the Backwards invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockBackwardsDone() bool {
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
func (m *CalculatorMock) MinimockBackwardsInspect() {
	for _, e := range m.BackwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.Backwards with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.BackwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		if m.BackwardsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.Backwards")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.Backwards with params: %#v", *m.BackwardsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBackwards != nil && mm_atomic.LoadUint64(&m.afterBackwardsCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.Backwards")
	}
}

type mCalculatorMockForwards struct {
	mock               *CalculatorMock
	defaultExpectation *CalculatorMockForwardsExpectation
	expectations       []*CalculatorMockForwardsExpectation

	callArgs []*CalculatorMockForwardsParams
	mutex    sync.RWMutex
}

// CalculatorMockForwardsExpectation specifies expectation struct of the Calculator.Forwards
type CalculatorMockForwardsExpectation struct {
	mock    *CalculatorMock
	params  *CalculatorMockForwardsParams
	results *CalculatorMockForwardsResults
	Counter uint64
}

// CalculatorMockForwardsParams contains parameters of the Calculator.Forwards
type CalculatorMockForwardsParams struct {
	ctx   context.Context
	pn    insolar.PulseNumber
	steps int
}

// CalculatorMockForwardsResults contains results of the Calculator.Forwards
type CalculatorMockForwardsResults struct {
	p1  insolar.Pulse
	err error
}

// Expect sets up expected params for Calculator.Forwards
func (mmForwards *mCalculatorMockForwards) Expect(ctx context.Context, pn insolar.PulseNumber, steps int) *mCalculatorMockForwards {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("CalculatorMock.Forwards mock is already set by Set")
	}

	if mmForwards.defaultExpectation == nil {
		mmForwards.defaultExpectation = &CalculatorMockForwardsExpectation{}
	}

	mmForwards.defaultExpectation.params = &CalculatorMockForwardsParams{ctx, pn, steps}
	for _, e := range mmForwards.expectations {
		if minimock.Equal(e.params, mmForwards.defaultExpectation.params) {
			mmForwards.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmForwards.defaultExpectation.params)
		}
	}

	return mmForwards
}

// Inspect accepts an inspector function that has same arguments as the Calculator.Forwards
func (mmForwards *mCalculatorMockForwards) Inspect(f func(ctx context.Context, pn insolar.PulseNumber, steps int)) *mCalculatorMockForwards {
	if mmForwards.mock.inspectFuncForwards != nil {
		mmForwards.mock.t.Fatalf("Inspect function is already set for CalculatorMock.Forwards")
	}

	mmForwards.mock.inspectFuncForwards = f

	return mmForwards
}

// Return sets up results that will be returned by Calculator.Forwards
func (mmForwards *mCalculatorMockForwards) Return(p1 insolar.Pulse, err error) *CalculatorMock {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("CalculatorMock.Forwards mock is already set by Set")
	}

	if mmForwards.defaultExpectation == nil {
		mmForwards.defaultExpectation = &CalculatorMockForwardsExpectation{mock: mmForwards.mock}
	}
	mmForwards.defaultExpectation.results = &CalculatorMockForwardsResults{p1, err}
	return mmForwards.mock
}

//Set uses given function f to mock the Calculator.Forwards method
func (mmForwards *mCalculatorMockForwards) Set(f func(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error)) *CalculatorMock {
	if mmForwards.defaultExpectation != nil {
		mmForwards.mock.t.Fatalf("Default expectation is already set for the Calculator.Forwards method")
	}

	if len(mmForwards.expectations) > 0 {
		mmForwards.mock.t.Fatalf("Some expectations are already set for the Calculator.Forwards method")
	}

	mmForwards.mock.funcForwards = f
	return mmForwards.mock
}

// When sets expectation for the Calculator.Forwards which will trigger the result defined by the following
// Then helper
func (mmForwards *mCalculatorMockForwards) When(ctx context.Context, pn insolar.PulseNumber, steps int) *CalculatorMockForwardsExpectation {
	if mmForwards.mock.funcForwards != nil {
		mmForwards.mock.t.Fatalf("CalculatorMock.Forwards mock is already set by Set")
	}

	expectation := &CalculatorMockForwardsExpectation{
		mock:   mmForwards.mock,
		params: &CalculatorMockForwardsParams{ctx, pn, steps},
	}
	mmForwards.expectations = append(mmForwards.expectations, expectation)
	return expectation
}

// Then sets up Calculator.Forwards return parameters for the expectation previously defined by the When method
func (e *CalculatorMockForwardsExpectation) Then(p1 insolar.Pulse, err error) *CalculatorMock {
	e.results = &CalculatorMockForwardsResults{p1, err}
	return e.mock
}

// Forwards implements pulse.Calculator
func (mmForwards *CalculatorMock) Forwards(ctx context.Context, pn insolar.PulseNumber, steps int) (p1 insolar.Pulse, err error) {
	mm_atomic.AddUint64(&mmForwards.beforeForwardsCounter, 1)
	defer mm_atomic.AddUint64(&mmForwards.afterForwardsCounter, 1)

	if mmForwards.inspectFuncForwards != nil {
		mmForwards.inspectFuncForwards(ctx, pn, steps)
	}

	params := &CalculatorMockForwardsParams{ctx, pn, steps}

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
		got := CalculatorMockForwardsParams{ctx, pn, steps}
		if want != nil && !minimock.Equal(*want, got) {
			mmForwards.t.Errorf("CalculatorMock.Forwards got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmForwards.ForwardsMock.defaultExpectation.results
		if results == nil {
			mmForwards.t.Fatal("No results are set for the CalculatorMock.Forwards")
		}
		return (*results).p1, (*results).err
	}
	if mmForwards.funcForwards != nil {
		return mmForwards.funcForwards(ctx, pn, steps)
	}
	mmForwards.t.Fatalf("Unexpected call to CalculatorMock.Forwards. %v %v %v", ctx, pn, steps)
	return
}

// ForwardsAfterCounter returns a count of finished CalculatorMock.Forwards invocations
func (mmForwards *CalculatorMock) ForwardsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForwards.afterForwardsCounter)
}

// ForwardsBeforeCounter returns a count of CalculatorMock.Forwards invocations
func (mmForwards *CalculatorMock) ForwardsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmForwards.beforeForwardsCounter)
}

// Calls returns a list of arguments used in each call to CalculatorMock.Forwards.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmForwards *mCalculatorMockForwards) Calls() []*CalculatorMockForwardsParams {
	mmForwards.mutex.RLock()

	argCopy := make([]*CalculatorMockForwardsParams, len(mmForwards.callArgs))
	copy(argCopy, mmForwards.callArgs)

	mmForwards.mutex.RUnlock()

	return argCopy
}

// MinimockForwardsDone returns true if the count of the Forwards invocations corresponds
// the number of defined expectations
func (m *CalculatorMock) MinimockForwardsDone() bool {
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
func (m *CalculatorMock) MinimockForwardsInspect() {
	for _, e := range m.ForwardsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to CalculatorMock.Forwards with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ForwardsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		if m.ForwardsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to CalculatorMock.Forwards")
		} else {
			m.t.Errorf("Expected call to CalculatorMock.Forwards with params: %#v", *m.ForwardsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcForwards != nil && mm_atomic.LoadUint64(&m.afterForwardsCounter) < 1 {
		m.t.Error("Expected call to CalculatorMock.Forwards")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *CalculatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockBackwardsInspect()

		m.MinimockForwardsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *CalculatorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *CalculatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockBackwardsDone() &&
		m.MinimockForwardsDone()
}
