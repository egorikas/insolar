package insolarapitests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed(t)
	fmt.Printf(seed)
	require.NotEmpty(t, seed)
}

func TestGetInfo(t *testing.T) {
	response := apihelper.GetInfo(t)
	require.NotEmpty(t, response.RootDomain)
	require.NotEmpty(t, response.RootMember)
	require.NotEmpty(t, response.NodeDomain)
	require.NotEmpty(t, response.TraceID)
}

func TestGetStatus(t *testing.T) {
	response := apihelper.GetStatus(t)
	require.NotEmpty(t, response.ActiveListSize)
}
