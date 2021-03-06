//
// Modified BSD 3-Clause Clear License
//
// Copyright (c) 2019 Insolar Technologies GmbH
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted (subject to the limitations in the disclaimer below) provided that
// the following conditions are met:
//  * Redistributions of source code must retain the above copyright notice, this list
//    of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright notice, this list
//    of conditions and the following disclaimer in the documentation and/or other materials
//    provided with the distribution.
//  * Neither the name of Insolar Technologies GmbH nor the names of its contributors
//    may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED
// BY THIS LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS
// AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
// INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY
// AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING,
// BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Notwithstanding any other provisions of this license, it is prohibited to:
//    (a) use this software,
//
//    (b) prepare modifications and derivative works of this software,
//
//    (c) distribute this software (including without limitation in source code, binary or
//        object code form), and
//
//    (d) reproduce copies of this software
//
//    for any commercial purposes, and/or
//
//    for the purposes of making available this software to third parties as a service,
//    including, without limitation, any software-as-a-service, platform-as-a-service,
//    infrastructure-as-a-service or other similar online service, irrespective of
//    whether it competes with the products or services of Insolar Technologies GmbH.
//

package power

import (
	"testing"

	"github.com/insolar/insolar/network/consensus/common/capacity"
	"github.com/insolar/insolar/network/consensus/gcpv2/api/member"
	"github.com/stretchr/testify/require"
)

func TestNewRequestByLevel(t *testing.T) {
	require.Equal(t, -Request(capacity.LevelMinimal)-1, NewRequestByLevel(capacity.LevelMinimal))
}

func TestNewRequest(t *testing.T) {
	require.Equal(t, Request(1)+1, NewRequest(member.Power(1)))
}

func TestAsCapacityLevel(t *testing.T) {
	b, l := Request(-1).AsCapacityLevel()
	require.True(t, b)
	require.Equal(t, capacity.Level(0), l)

	b, l = Request(1).AsCapacityLevel()
	require.False(t, b)

	r := Request(-2)
	require.Equal(t, capacity.Level(r), l)

	b, l = Request(0).AsCapacityLevel()
	require.False(t, b)

	r = Request(-1)
	require.Equal(t, capacity.Level(r), l)
}

func TestAsMemberPower(t *testing.T) {
	b, l := Request(1).AsMemberPower()
	require.True(t, b)
	require.Zero(t, l)

	b, l = Request(-1).AsMemberPower()
	require.False(t, b)

	r := Request(-2)
	require.Equal(t, member.Power(r), l)

	b, l = Request(0).AsMemberPower()
	require.False(t, b)

	r = Request(-1)
	require.Equal(t, member.Power(r), l)
}

func TestIsEmpty(t *testing.T) {
	require.True(t, EmptyRequest.IsEmpty())

	require.False(t, Request(1).IsEmpty())
}

func TestUpdate(t *testing.T) {
	pws := member.PowerSet([...]member.Power{10, 20, 30, 40})
	pwBase := member.Power(1)
	pw := pwBase

	require.True(t, Request(-1).Update(&pw, pws))

	require.Zero(t, pw)

	pw = pwBase
	require.True(t, Request(-2).Update(&pw, pws))

	require.Equal(t, member.Power(10), pw)

	pw = pwBase
	require.True(t, Request(10).Update(&pw, pws))

	require.Equal(t, member.Power(10), pw)

	pw = pwBase
	require.True(t, Request(100).Update(&pw, pws))

	require.Equal(t, member.Power(40), pw)

	pw = pwBase
	require.False(t, Request(0).Update(&pw, pws))

	require.Equal(t, member.Power(1), pw)
}
