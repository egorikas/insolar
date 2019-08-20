/*
 * Copyright 2019 Insolar Technologies GmbH
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package executor

import (
	"testing"

	"github.com/gojuno/minimock"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/gen"
	"github.com/insolar/insolar/insolar/jet"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/ledger/drop"
	"github.com/insolar/insolar/ledger/object"
	"github.com/stretchr/testify/require"
)

var (
	pending insolar.PulseNumber = 780
	topSync insolar.PulseNumber = 800
	current insolar.PulseNumber = 850
)

func indexesFixture() []record.Index {
	objIDs := gen.UniqueIDs(5)
	return []record.Index{
		{
			ObjID: objIDs[0],
			Lifeline: record.Lifeline{
				EarliestOpenRequest: &pending,
			},
		},
		{
			ObjID: objIDs[1],
			Lifeline: record.Lifeline{
				EarliestOpenRequest: &pending,
			},
		},
		{
			ObjID: objIDs[2],
			Lifeline: record.Lifeline{
				EarliestOpenRequest: &pending,
			},
		},
		{
			ObjID: objIDs[3],
			Lifeline: record.Lifeline{
				EarliestOpenRequest: nil,
			},
		},
		{
			ObjID: objIDs[4],
			Lifeline: record.Lifeline{
				EarliestOpenRequest: nil,
			},
		},
	}
}

func dropsFixture() []drop.Drop {
	ids := gen.UniqueIDs(3)
	return []drop.Drop{
		{Split: false, Hash: ids[0].Hash()},
		{Split: true, Hash: ids[1].Hash()},
		{Split: false, Hash: ids[2].Hash()},
	}
}

func TestInitialStateKeeper_Get_AfterRestart(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()
	ctx := inslogger.TestContext(t)

	topSyncPulse := insolar.Pulse{PulseNumber: topSync}
	jetKeeper := NewJetKeeperMock(mc)
	jetKeeper.TopSyncPulseMock.Return(topSyncPulse.PulseNumber)

	jetIDs := gen.UniqueJetIDs(3)
	// split jet
	left, right := jet.Siblings(jetIDs[1])

	jetAccessor := jet.NewAccessorMock(mc)
	jetAccessor.AllMock.Expect(ctx, topSyncPulse.PulseNumber).Return(jetIDs)

	drops := dropsFixture()
	dropAccessor := drop.NewAccessorMock(mc)
	dropAccessor.ForPulseMock.When(ctx, jetIDs[0], topSyncPulse.PulseNumber).Then(drops[0], nil)
	dropAccessor.ForPulseMock.When(ctx, jetIDs[1], topSyncPulse.PulseNumber).Then(drops[1], nil)
	dropAccessor.ForPulseMock.When(ctx, jetIDs[2], topSyncPulse.PulseNumber).Then(drops[2], nil)

	indexes := indexesFixture()
	indexAccessor := object.NewIndexAccessorMock(mc)
	indexAccessor.ForPulseMock.Expect(ctx, topSyncPulse.PulseNumber).Return(indexes, nil)

	jetAccessor.ForIDMock.When(ctx, topSync, indexes[0].ObjID).Then(jetIDs[0], true)
	jetAccessor.ForIDMock.When(ctx, topSync, indexes[1].ObjID).Then(jetIDs[0], true)
	jetAccessor.ForIDMock.When(ctx, topSync, indexes[2].ObjID).Then(jetIDs[2], true)

	jetCoordinator := jet.NewCoordinatorMock(mc)

	stateKeeper := NewInitialStateKeeper(jetKeeper, jetAccessor, jetCoordinator, indexAccessor, dropAccessor)
	err := stateKeeper.Start(ctx)
	require.NoError(t, err)

	currentLight := gen.Reference()
	anotherLight := gen.Reference()

	jetCoordinator.LightExecutorForJetMock.When(ctx, insolar.ID(jetIDs[0]), current).Then(&currentLight, nil)
	jetCoordinator.LightExecutorForJetMock.When(ctx, insolar.ID(left), current).Then(&currentLight, nil)
	jetCoordinator.LightExecutorForJetMock.When(ctx, insolar.ID(right), current).Then(&anotherLight, nil)
	jetCoordinator.LightExecutorForJetMock.When(ctx, insolar.ID(jetIDs[2]), current).Then(&anotherLight, nil)

	state := stateKeeper.Get(ctx, currentLight, current)
	require.Equal(t, []record.Index{indexes[0], indexes[1]}, state.Indexes)
	require.Equal(t, [][]byte{
		drop.MustEncode(&drops[0]),
		drop.MustEncode(&drops[1]),
	}, state.Drops)
	require.Equal(t, []insolar.JetID{jetIDs[0], left}, state.JetIDs)
}