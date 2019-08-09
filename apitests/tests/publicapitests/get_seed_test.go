package publicapitests

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed(t)
	require.NotEmpty(t, seed)
}

type cases struct {
	input         string
	expectedError error
}

type casesInt struct {
	input         int32
	expectedError error
}

type error struct {
	Code    int
	Message string
}

func TestGetSeedWithBadMethod(t *testing.T) {
	randomString := testutils.RandomString()
	data := []cases{
		{"", error{-32000, "rpc: service/method request ill-formed: \"\""}},
		{" ", error{-32000, "rpc: service/method request ill-formed: \" \""}},
		{"node.getInfo", error{-32000, "rpc: can't find method \"node.getInfo\""}},
		{randomString, error{-32000, "rpc: service/method request ill-formed: \"" + randomString + "\""}},
		{"1111", error{-32000, "rpc: service/method request ill-formed: \"1111\""}},
		{"node getSeed", error{-32000, "rpc: service/method request ill-formed: \"node getSeed\""}},
		{"node^getSeed", error{-32000, "rpc: service/method request ill-formed: \"node^getSeed\""}},
		{"node*getSeed", error{-32000, "rpc: service/method request ill-formed: \"node*getSeed\""}},
		{"node&getSeed", error{-32000, "rpc: service/method request ill-formed: \"node&getSeed\""}},
		{"node%getSeed", error{-32000, "rpc: service/method request ill-formed: \"node%getSeed\""}},
		{"getSeed", error{-32000, "rpc: service/method request ill-formed: \"getSeed\""}},
	}
	for _, tc := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  tc.input,
		}
		getSeedWithBadRequest(t, r, tc.expectedError)
	}
}

func TestGetSeedWithoutMethod(t *testing.T) {
	r := insolar_api.NodeGetSeedRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
	}
	getSeedWithBadRequest(t, r, error{-32000, "rpc: service/method request ill-formed: \"\""}) //todo bug - need error: "method name is required"
}

func TestGetSeedWithBadJsonVersion(t *testing.T) {
	randomString := testutils.RandomString()
	data := []cases{
		{"1.0", error{-32600, "jsonrpc must be 2.0"}},
		{"", error{-32600, "jsonrpc must be 2.0"}},
		{" ", error{-32600, "jsonrpc must be 2.0"}},
		{randomString, error{-32600, "jsonrpc must be 2.0"}},
		{"0", error{-32600, "jsonrpc must be 2.0"}},
		{"-1", error{-32600, "jsonrpc must be 2.0"}},
		{"3.0", error{-32600, "jsonrpc must be 2.0"}},
	}
	for _, tc := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: tc.input,
			Id:      apihelper.GetRequestId(),
			Method:  apihelper.GetSeedMethod,
		}
		getSeedWithBadRequest(t, r, tc.expectedError)
	}
}

func TestGetSeedWithoutJsonField(t *testing.T) {
	r := insolar_api.NodeGetSeedRequest{
		Id:     apihelper.GetRequestId(),
		Method: apihelper.GetSeedMethod,
	}
	getSeedWithBadRequest(t, r, error{-32600, "jsonrpc must be 2.0"})
}

func TestGetSeedWithBadRequestId(t *testing.T) {
	data := []casesInt{
		//{0, error{-32600,"jsonrpc must be 2.0"},},//todo какие требования к id?
		//{-1, error{-32600,"jsonrpc must be 2.0"},},//todo какие требования к id?
		{999000000, error{-32600, "jsonrpc must be 2.0"}},
	}
	for _, tc := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      tc.input,
			Method:  apihelper.GetSeedMethod,
		}
		getSeedWithBadRequest(t, r, tc.expectedError)
	}
}

func getSeedWithBadRequest(t *testing.T, r insolar_api.NodeGetSeedRequest, error error) {
	apihelper.Logger.Printf("%v request body:\n %v", apihelper.GetSeedMethod, r)
	response, http, err := apihelper.GetClient().InformationApi.GetSeed(nil, r)
	apihelper.Logger.Printf("%v response statusCode:\n %v", apihelper.GetSeedMethod, http.StatusCode)
	apihelper.Logger.Printf("%v response id:\n %v", apihelper.GetSeedMethod, response.Id)
	apihelper.Logger.Printf("%v response body:\n %v", apihelper.GetSeedMethod, response)
	apihelper.Logger.Printf("%v response Err:\n %v", apihelper.GetSeedMethod, response.Error)
	if err != nil {
		log.Fatalln(err)
	}
	require.Equal(t, 200, http.StatusCode)
	require.Empty(t, response.Result)
	require.Equal(t, error.Message, response.Error.Message)
	require.Equal(t, int32(error.Code), response.Error.Code)
}
