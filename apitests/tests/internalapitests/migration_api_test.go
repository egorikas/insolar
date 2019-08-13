package internalapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrationAddAddresses(t *testing.T) {
	response := apihelper.AddMigrationAddresses(t)
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestMigrationDeposit(t *testing.T) {
	response := apihelper.MigrationDeposit(t)
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}

func TestObserverGetToken(t *testing.T) {
	response := apihelper.ObserverToken(t) //not worked
	require.NotEmpty(t, response)
}

func TestMemberGetBalance(t *testing.T) {
	member := apihelper.CreateMember(t)
	response := apihelper.GetBalance(t, member)
	require.NotEmpty(t, response.Result.CallResult.Deposits)
	require.NotEmpty(t, response.Result.CallResult.Balance)
}

/* "code": 217,
   "message": "[ makeCall ] Error in called method: unknown method: 'member.getBalance'"*/

func TestMigrationDeactivateDaemon(t *testing.T) {
	response := apihelper.MigrationDeactivateDaemon(t, "")
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
}
