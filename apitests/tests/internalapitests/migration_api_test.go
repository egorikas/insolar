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
