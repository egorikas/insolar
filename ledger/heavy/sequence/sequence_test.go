//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sequence

import (
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/record"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/internal/ledger/store"
	"github.com/insolar/insolar/ledger/blob"
	"github.com/insolar/insolar/ledger/object"
)

func TestSequencer_FromRecords(t *testing.T) {
	ctx := inslogger.TestContext(t)

	db := store.NewMemoryMockDB()
	storage := object.NewRecordDB(db)
	records := NewSequencer(db, store.ScopeRecord)

	type tempRecord struct {
		id  insolar.ID
		rec record.Material
	}

	const (
		pulse = 10
		index = 10
		limit = 10
		total = 30
	)

	var all []tempRecord

	f := fuzz.New().Funcs(func(t *tempRecord, c fuzz.Continue) {
		var hash [insolar.RecordIDSize - insolar.RecordHashOffset]byte
		c.Fuzz(&hash)
		t.id = *insolar.NewID(pulse, hash[:])
		t.rec = record.Material{Polymorph: c.Int31()}
	})
	f.NilChance(0)
	f.NumElements(30, 30)
	f.Fuzz(&all)

	for i := 0; i < total; i++ {
		err := storage.Set(ctx, all[i].id, all[i].rec)
		require.NoError(t, err)
	}

	recs, err := records.From(ctx, pulse, index, limit)
	require.NoError(t, err)
	require.Equal(t, limit, len(recs))
}

func TestSequencer_FromBlobs(t *testing.T) {
	ctx := inslogger.TestContext(t)

	db := store.NewMemoryMockDB()
	storage := blob.NewDB(db)
	records := NewSequencer(db, store.ScopeBlob)

	type tempBlob struct {
		id   insolar.ID
		blob blob.Blob
	}

	const (
		pulse = 10
		index = 10
		limit = 10
		total = 30
	)

	var all []tempBlob

	f := fuzz.New().Funcs(func(t *tempBlob, c fuzz.Continue) {
		var memory [10]byte
		c.Fuzz(&memory)
		var hash [insolar.RecordIDSize - insolar.RecordHashOffset]byte
		c.Fuzz(&hash)
		t.id = *insolar.NewID(pulse, hash[:])
		t.blob = blob.Blob{
			JetID: insolar.ZeroJetID,
			Value: memory[:],
		}
	})
	f.NilChance(0)
	f.NumElements(30, 30)
	f.Fuzz(&all)

	for i := 0; i < total; i++ {
		err := storage.Set(ctx, all[i].id, all[i].blob)
		require.NoError(t, err)
	}

	recs, err := records.From(ctx, pulse, index, limit)
	require.NoError(t, err)
	require.Equal(t, limit, len(recs))
}
