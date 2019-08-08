package publicapitests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

//getSeed Tests----------------------------------------------------------------------------------------------

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed(t)
	fmt.Printf(seed)
	require.NotEmpty(t, seed)
}

func TestGetSeedWithBadMethod(t *testing.T) {
	//r := insolar_api.NodeGetSeedRequest{
	//	Jsonrpc: apihelper.JSONRPCVersion,
	//	Id:      apihelper.GetRequestId(),
	//	Method:  "node.getInfo",
	//}
	//responseData, httpResponse, err := apihelper.GetClient().InformationApi.GetSeed(nil, r)
	//fmt.Printf(json.Marshal(responseData))
	//fmt.Printf(httpResponse)
	//fmt.Printf(err)
	//require.NotEmpty(t, responseData)
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
