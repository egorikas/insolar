package publicapitests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetStatus(t *testing.T) {
	response := apihelper.GetStatus(t)
	require.NotEmpty(t, response.ActiveListSize)
}
