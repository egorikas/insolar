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

// BackupMakerMock implements executor.BackupMaker
type BackupMakerMock struct {
	t minimock.Tester

	funcMakeBackup          func(ctx context.Context, lastFinalizedPulse insolar.PulseNumber) (err error)
	inspectFuncMakeBackup   func(ctx context.Context, lastFinalizedPulse insolar.PulseNumber)
	afterMakeBackupCounter  uint64
	beforeMakeBackupCounter uint64
	MakeBackupMock          mBackupMakerMockMakeBackup
}

// NewBackupMakerMock returns a mock for executor.BackupMaker
func NewBackupMakerMock(t minimock.Tester) *BackupMakerMock {
	m := &BackupMakerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.MakeBackupMock = mBackupMakerMockMakeBackup{mock: m}
	m.MakeBackupMock.callArgs = []*BackupMakerMockMakeBackupParams{}

	return m
}

type mBackupMakerMockMakeBackup struct {
	mock               *BackupMakerMock
	defaultExpectation *BackupMakerMockMakeBackupExpectation
	expectations       []*BackupMakerMockMakeBackupExpectation

	callArgs []*BackupMakerMockMakeBackupParams
	mutex    sync.RWMutex
}

// BackupMakerMockMakeBackupExpectation specifies expectation struct of the BackupMaker.MakeBackup
type BackupMakerMockMakeBackupExpectation struct {
	mock    *BackupMakerMock
	params  *BackupMakerMockMakeBackupParams
	results *BackupMakerMockMakeBackupResults
	Counter uint64
}

// BackupMakerMockMakeBackupParams contains parameters of the BackupMaker.MakeBackup
type BackupMakerMockMakeBackupParams struct {
	ctx                context.Context
	lastFinalizedPulse insolar.PulseNumber
}

// BackupMakerMockMakeBackupResults contains results of the BackupMaker.MakeBackup
type BackupMakerMockMakeBackupResults struct {
	err error
}

// Expect sets up expected params for BackupMaker.MakeBackup
func (mmMakeBackup *mBackupMakerMockMakeBackup) Expect(ctx context.Context, lastFinalizedPulse insolar.PulseNumber) *mBackupMakerMockMakeBackup {
	if mmMakeBackup.mock.funcMakeBackup != nil {
		mmMakeBackup.mock.t.Fatalf("BackupMakerMock.MakeBackup mock is already set by Set")
	}

	if mmMakeBackup.defaultExpectation == nil {
		mmMakeBackup.defaultExpectation = &BackupMakerMockMakeBackupExpectation{}
	}

	mmMakeBackup.defaultExpectation.params = &BackupMakerMockMakeBackupParams{ctx, lastFinalizedPulse}
	for _, e := range mmMakeBackup.expectations {
		if minimock.Equal(e.params, mmMakeBackup.defaultExpectation.params) {
			mmMakeBackup.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmMakeBackup.defaultExpectation.params)
		}
	}

	return mmMakeBackup
}

// Inspect accepts an inspector function that has same arguments as the BackupMaker.MakeBackup
func (mmMakeBackup *mBackupMakerMockMakeBackup) Inspect(f func(ctx context.Context, lastFinalizedPulse insolar.PulseNumber)) *mBackupMakerMockMakeBackup {
	if mmMakeBackup.mock.inspectFuncMakeBackup != nil {
		mmMakeBackup.mock.t.Fatalf("Inspect function is already set for BackupMakerMock.MakeBackup")
	}

	mmMakeBackup.mock.inspectFuncMakeBackup = f

	return mmMakeBackup
}

// Return sets up results that will be returned by BackupMaker.MakeBackup
func (mmMakeBackup *mBackupMakerMockMakeBackup) Return(err error) *BackupMakerMock {
	if mmMakeBackup.mock.funcMakeBackup != nil {
		mmMakeBackup.mock.t.Fatalf("BackupMakerMock.MakeBackup mock is already set by Set")
	}

	if mmMakeBackup.defaultExpectation == nil {
		mmMakeBackup.defaultExpectation = &BackupMakerMockMakeBackupExpectation{mock: mmMakeBackup.mock}
	}
	mmMakeBackup.defaultExpectation.results = &BackupMakerMockMakeBackupResults{err}
	return mmMakeBackup.mock
}

//Set uses given function f to mock the BackupMaker.MakeBackup method
func (mmMakeBackup *mBackupMakerMockMakeBackup) Set(f func(ctx context.Context, lastFinalizedPulse insolar.PulseNumber) (err error)) *BackupMakerMock {
	if mmMakeBackup.defaultExpectation != nil {
		mmMakeBackup.mock.t.Fatalf("Default expectation is already set for the BackupMaker.MakeBackup method")
	}

	if len(mmMakeBackup.expectations) > 0 {
		mmMakeBackup.mock.t.Fatalf("Some expectations are already set for the BackupMaker.MakeBackup method")
	}

	mmMakeBackup.mock.funcMakeBackup = f
	return mmMakeBackup.mock
}

// When sets expectation for the BackupMaker.MakeBackup which will trigger the result defined by the following
// Then helper
func (mmMakeBackup *mBackupMakerMockMakeBackup) When(ctx context.Context, lastFinalizedPulse insolar.PulseNumber) *BackupMakerMockMakeBackupExpectation {
	if mmMakeBackup.mock.funcMakeBackup != nil {
		mmMakeBackup.mock.t.Fatalf("BackupMakerMock.MakeBackup mock is already set by Set")
	}

	expectation := &BackupMakerMockMakeBackupExpectation{
		mock:   mmMakeBackup.mock,
		params: &BackupMakerMockMakeBackupParams{ctx, lastFinalizedPulse},
	}
	mmMakeBackup.expectations = append(mmMakeBackup.expectations, expectation)
	return expectation
}

// Then sets up BackupMaker.MakeBackup return parameters for the expectation previously defined by the When method
func (e *BackupMakerMockMakeBackupExpectation) Then(err error) *BackupMakerMock {
	e.results = &BackupMakerMockMakeBackupResults{err}
	return e.mock
}

// MakeBackup implements executor.BackupMaker
func (mmMakeBackup *BackupMakerMock) MakeBackup(ctx context.Context, lastFinalizedPulse insolar.PulseNumber) (err error) {
	mm_atomic.AddUint64(&mmMakeBackup.beforeMakeBackupCounter, 1)
	defer mm_atomic.AddUint64(&mmMakeBackup.afterMakeBackupCounter, 1)

	if mmMakeBackup.inspectFuncMakeBackup != nil {
		mmMakeBackup.inspectFuncMakeBackup(ctx, lastFinalizedPulse)
	}

	params := &BackupMakerMockMakeBackupParams{ctx, lastFinalizedPulse}

	// Record call args
	mmMakeBackup.MakeBackupMock.mutex.Lock()
	mmMakeBackup.MakeBackupMock.callArgs = append(mmMakeBackup.MakeBackupMock.callArgs, params)
	mmMakeBackup.MakeBackupMock.mutex.Unlock()

	for _, e := range mmMakeBackup.MakeBackupMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmMakeBackup.MakeBackupMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMakeBackup.MakeBackupMock.defaultExpectation.Counter, 1)
		want := mmMakeBackup.MakeBackupMock.defaultExpectation.params
		got := BackupMakerMockMakeBackupParams{ctx, lastFinalizedPulse}
		if want != nil && !minimock.Equal(*want, got) {
			mmMakeBackup.t.Errorf("BackupMakerMock.MakeBackup got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmMakeBackup.MakeBackupMock.defaultExpectation.results
		if results == nil {
			mmMakeBackup.t.Fatal("No results are set for the BackupMakerMock.MakeBackup")
		}
		return (*results).err
	}
	if mmMakeBackup.funcMakeBackup != nil {
		return mmMakeBackup.funcMakeBackup(ctx, lastFinalizedPulse)
	}
	mmMakeBackup.t.Fatalf("Unexpected call to BackupMakerMock.MakeBackup. %v %v", ctx, lastFinalizedPulse)
	return
}

// MakeBackupAfterCounter returns a count of finished BackupMakerMock.MakeBackup invocations
func (mmMakeBackup *BackupMakerMock) MakeBackupAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMakeBackup.afterMakeBackupCounter)
}

// MakeBackupBeforeCounter returns a count of BackupMakerMock.MakeBackup invocations
func (mmMakeBackup *BackupMakerMock) MakeBackupBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMakeBackup.beforeMakeBackupCounter)
}

// Calls returns a list of arguments used in each call to BackupMakerMock.MakeBackup.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmMakeBackup *mBackupMakerMockMakeBackup) Calls() []*BackupMakerMockMakeBackupParams {
	mmMakeBackup.mutex.RLock()

	argCopy := make([]*BackupMakerMockMakeBackupParams, len(mmMakeBackup.callArgs))
	copy(argCopy, mmMakeBackup.callArgs)

	mmMakeBackup.mutex.RUnlock()

	return argCopy
}

// MinimockMakeBackupDone returns true if the count of the MakeBackup invocations corresponds
// the number of defined expectations
func (m *BackupMakerMock) MinimockMakeBackupDone() bool {
	for _, e := range m.MakeBackupMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MakeBackupMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMakeBackupCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMakeBackup != nil && mm_atomic.LoadUint64(&m.afterMakeBackupCounter) < 1 {
		return false
	}
	return true
}

// MinimockMakeBackupInspect logs each unmet expectation
func (m *BackupMakerMock) MinimockMakeBackupInspect() {
	for _, e := range m.MakeBackupMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BackupMakerMock.MakeBackup with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MakeBackupMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMakeBackupCounter) < 1 {
		if m.MakeBackupMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BackupMakerMock.MakeBackup")
		} else {
			m.t.Errorf("Expected call to BackupMakerMock.MakeBackup with params: %#v", *m.MakeBackupMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMakeBackup != nil && mm_atomic.LoadUint64(&m.afterMakeBackupCounter) < 1 {
		m.t.Error("Expected call to BackupMakerMock.MakeBackup")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BackupMakerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockMakeBackupInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BackupMakerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *BackupMakerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockMakeBackupDone()
}
