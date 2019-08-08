package publicapitests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

//getSeed Tests----------------------------------------------------------------------------------------------

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed(t)
	fmt.Printf(seed)
	require.NotEmpty(t, seed)
}

func TestGetSeedWithBadMethod(t *testing.T) {
	r := insolar_api.NodeGetSeedRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
		Method:  "node.getInfo",
	}
	response, http, err := apihelper.GetClient().InformationApi.GetSeed(nil, r)
	apihelper.Logger.Printf("%v response statusCode:\n %v", apihelper.GETSEED, http.StatusCode)
	apihelper.Logger.Printf("%v response id:\n %v", apihelper.GETSEED, response.Id)
	apihelper.Logger.Printf("%v response body:\n %v", apihelper.GETSEED, response)
	apihelper.Logger.Printf("%v response Err:\n %v", apihelper.GETSEED, response.Error)
	if err != nil {
		log.Fatalln(err)
	}
	require.Equal(t, 200, http.StatusCode)
	require.Empty(t, response.Result)
	require.Equal(t, "rpc: can't find method \"node.getInfo\"", response.Error.Message)
	require.Equal(t, int32(-32000), response.Error.Code)
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
