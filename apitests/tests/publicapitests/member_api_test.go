package publicapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMember(t *testing.T) {
	member := apihelper.CreateMember(t)
	require.NotEmpty(t, member.MemberReference, "MemberReference")
}

func TestMemberTransfer(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	member2 := apihelper.CreateMember(t)
	transfer := member1.Transfer(t, member2.MemberReference, "1")
	require.NotEmpty(t, transfer.Result.CallResult.Fee, "Fee")
}

func TestGetMember(t *testing.T) {
	member1 := apihelper.CreateMember(t)
	resp := member1.GetMember(t)
	require.Equal(t, member1.MemberReference, resp.Result.CallResult.Reference, "Reference")
	require.Empty(t, resp.Result.CallResult.MigrationAddress, "MigrationAddress")
}
