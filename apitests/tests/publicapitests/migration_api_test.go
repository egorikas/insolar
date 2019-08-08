package publicapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositTransfer(t *testing.T) {
	response := apihelper.DepositTransfer()
	require.NotEmpty(t, response.Result.CallResult)
}

func TestMemberMigrationCreate(t *testing.T) {
	var member = apihelper.MemberMigrationCreate()
	require.NotEmpty(t, member)
	require.NotEmpty(t, member.MemberResponseResult)
}
