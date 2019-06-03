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
	"context"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/internal/ledger/store"
)

type Item struct {
	Index  uint32
	ID     insolar.ID
	Record []byte
}

// Sequencer is an interface to work with records sequence.
type Sequencer interface {
	// From returns slice from provided position with corresponding limit.
	From(ctx context.Context, pulse insolar.PulseNumber, index, limit uint32) ([]Item, error)
}

type sequencer struct {
	db    store.DB
	scope store.Scope
}

func NewSequencer(db store.DB, scope store.Scope) Sequencer {
	return &sequencer{db: db, scope: scope}
}

func (s *sequencer) From(ctx context.Context, pulse insolar.PulseNumber, index, limit uint32) ([]Item, error) {
	var result []Item
	it := s.db.NewIterator(s.scope)
	it.Seek(pulse.Bytes())
	for i := uint32(0); it.Next(); i++ {
		if i >= index && len(result) < int(limit) {
			buff, err := it.Value()
			if err != nil {
				return nil, errors.Wrapf(err, "failed to get Value from db iterator")
			}
			key := it.Key()
			id := [32]byte{}
			copy(id[:], key)
			result = append(result, Item{
				Index:  i,
				ID:     id,
				Record: buff,
			})
		}
		if len(result) == int(limit) {
			break
		}
	}
	return result, nil
}
