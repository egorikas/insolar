package publicapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateMember(t *testing.T) {
	member := apihelper.CreateMember()
	require.NotEmpty(t, member.MemberResponseResult.Result.CallResult.Reference)
}
