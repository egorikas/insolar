package publicapitests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed()
	fmt.Printf(seed)
	require.NotEmpty(t, seed)
}

func TestGetInfo(t *testing.T) {
	response := apihelper.GetInfo()
	require.NotEqual(t, "", response.RootDomain)
	require.NotEqual(t, "", response.RootMember)
	require.NotEqual(t, "", response.NodeDomain)
	require.NotEqual(t, "", response.TraceID)
}

func TestGetStatus(t *testing.T) {
	response := apihelper.GetStatus()
	require.NotEmpty(t, response)
}
