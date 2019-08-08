package publicapitests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"testing"
)

//getSeed Tests----------------------------------------------------------------------------------------------

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed()
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
