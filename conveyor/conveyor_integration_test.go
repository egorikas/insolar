/*
 *    Copyright 2019 Insolar Technologies
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package conveyor

import (
	"context"
	"encoding/hex"
	"os"
	"testing"
	"time"

	"github.com/insolar/insolar/component"
	"github.com/insolar/insolar/conveyor/adapter/adapterstorage"
	"github.com/insolar/insolar/conveyor/fsm"
	"github.com/insolar/insolar/conveyor/generator/matrix"
	"github.com/insolar/insolar/conveyor/handler"
	"github.com/insolar/insolar/conveyor/slot"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/reply"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
)

const maxState = fsm.StateID(1000)

type mockStateMachineSet struct {
	stateMachine matrix.StateMachine
}

func (s *mockStateMachineSet) GetStateMachineByID(id int) matrix.StateMachine {
	return s.stateMachine
}

type mockStateMachineHolder struct{}

func (m *mockStateMachineHolder) makeSetAccessor() matrix.SetAccessor {
	return &mockStateMachineSet{
		stateMachine: m.GetStateMachinesByType(),
	}
}

func (m *mockStateMachineHolder) GetFutureConfig() matrix.SetAccessor {
	return m.makeSetAccessor()
}

func (m *mockStateMachineHolder) GetPresentConfig() matrix.SetAccessor {
	return m.makeSetAccessor()
}

func (m *mockStateMachineHolder) GetPastConfig() matrix.SetAccessor {
	return m.makeSetAccessor()
}

func (m *mockStateMachineHolder) GetInitialStateMachine() matrix.StateMachine {
	return m.GetStateMachinesByType()
}

func (m *mockStateMachineHolder) GetStateMachinesByType() matrix.StateMachine {

	sm := matrix.NewStateMachineMock(&testing.T{})
	sm.GetMigrationHandlerFunc = func(s fsm.StateID) (r handler.MigrationHandler) {
		return func(element fsm.SlotElementHelper) (interface{}, fsm.ElementState, error) {
			if s > maxState {
				s /= 2
			}
			return element.GetElementID(), fsm.NewElementState(fsm.ID(s%3), s+1), nil
		}
	}

	sm.GetTransitionHandlerFunc = func(s fsm.StateID) (r handler.TransitHandler) {
		return func(element fsm.SlotElementHelper) (interface{}, fsm.ElementState, error) {
			if s > maxState {
				s /= 2
			}
			return element.GetElementID(), fsm.NewElementState(fsm.ID(s%3), s+1), nil
		}
	}

	sm.GetResponseHandlerFunc = func(s fsm.StateID) (r handler.AdapterResponseHandler) {
		return func(element fsm.SlotElementHelper, response interface{}) (interface{}, fsm.ElementState, error) {
			if s > maxState {
				s /= 2
			}
			return element.GetPayload(), fsm.NewElementState(fsm.ID(s%3), s+1), nil
		}
	}

	return sm
}

func mockHandlerStorage() matrix.StateMachineHolder {
	return &mockStateMachineHolder{}
}

func initComponents(t *testing.T) {
	pc := testutils.NewPlatformCryptographyScheme()
	ledgerMock := testutils.NewLedgerLogicMock(t)
	ledgerMock.GetCodeFunc = func(p context.Context, p1 insolar.Parcel) (r insolar.Reply, r1 error) {
		return &reply.Code{}, nil
	}

	cm := &component.Manager{}
	ctx := context.TODO()

	components := adapterstorage.GetAllProcessors()
	components = append(components, pc, ledgerMock)
	cm.Inject(components...)
	err := cm.Init(ctx)
	if err != nil {
		t.Error("ComponentManager init failed", err)
	}
}

func setup() {
	slot.HandlerStorage = mockHandlerStorage()
}

func testMainWrapper(m *testing.M) int {
	setup()
	code := m.Run()
	return code
}

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func TestConveyor_ChangePulse(t *testing.T) {
	conveyor, err := NewPulseConveyor()
	require.NoError(t, err)
	initComponents(t)
	callback := mockCallback()
	pulse := insolar.Pulse{PulseNumber: testRealPulse + testPulseDelta}
	err = conveyor.PreparePulse(pulse, callback)
	require.NoError(t, err)

	callback.(*mockSyncDone).GetResult()

	err = conveyor.ActivatePulse()
	require.NoError(t, err)
}

func TestConveyor_ChangePulseMultipleTimes(t *testing.T) {
	conveyor, err := NewPulseConveyor()
	require.NoError(t, err)
	initComponents(t)

	pulseNumber := testRealPulse + testPulseDelta
	for i := 0; i < 20; i++ {
		callback := mockCallback()
		pulseNumber += testPulseDelta
		pulse := insolar.Pulse{PulseNumber: pulseNumber, NextPulseNumber: pulseNumber + testPulseDelta}
		err = conveyor.PreparePulse(pulse, callback)
		require.NoError(t, err)

		callback.(*mockSyncDone).GetResult()

		err = conveyor.ActivatePulse()
		require.NoError(t, err)
	}
}

func TestConveyor_ChangePulseMultipleTimes_WithEvents(t *testing.T) {
	conveyor, err := NewPulseConveyor()
	require.NoError(t, err)
	initComponents(t)

	pulseNumber := testRealPulse + testPulseDelta
	for i := 0; i < 100; i++ {

		go func() {
			for j := 0; j < 1; j++ {
				conveyor.SinkPush(pulseNumber, "TEST")
				conveyor.SinkPush(pulseNumber-testPulseDelta, "TEST")
				conveyor.SinkPush(pulseNumber+testPulseDelta, "TEST")
				conveyor.SinkPushAll(pulseNumber, []interface{}{"TEST", i * j})
			}
		}()

		go func() {
			for j := 0; j < 100; j++ {
				conveyor.GetState()
			}
		}()

		go func() {
			for j := 0; j < 100; j++ {
				conveyor.IsOperational()
			}
		}()

		callback := mockCallback()
		pulseNumber += testPulseDelta
		pulse := insolar.Pulse{PulseNumber: pulseNumber, NextPulseNumber: pulseNumber + testPulseDelta}
		err = conveyor.PreparePulse(pulse, callback)
		require.NoError(t, err)

		expectedHash, _ := hex.DecodeString(
			"0c60ae04fbb17fe36f4e84631a5b8f3cd6d0cd46e80056bdfec97fd305f764daadef8ae1adc89b203043d7e2af1fb341df0ce5f66dfe3204ec3a9831532a8e4c",
		)
		require.Equal(t, expectedHash, callback.(*mockSyncDone).GetResult())

		err = conveyor.ActivatePulse()
		require.NoError(t, err)

		go func() {
			for j := 0; j < 10; j++ {
				require.NoError(t, conveyor.SinkPushAll(pulseNumber, []interface{}{"TEST", i}))
				require.NoError(t, conveyor.SinkPush(pulseNumber, "TEST"))
				require.NoError(t, conveyor.SinkPush(pulseNumber-testPulseDelta, "TEST"))
				conveyor.SinkPush(pulseNumber+testPulseDelta, "TEST")
			}
		}()
	}

	time.Sleep(time.Millisecond * 200)
}

// TODO: Add test on InitiateShutdown
