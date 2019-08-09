package publicapitests

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	"github.com/insolar/insolar/apitests/tests"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGetSeedWithBadMethod(t *testing.T) {
	randomString := testutils.RandomString()
	data := []tests.Cases{
		{"", tests.TestError{-32000, "rpc: service/method request ill-formed: \"\""}},
		{" ", tests.TestError{-32000, "rpc: service/method request ill-formed: \" \""}},
		{"node.getInfo", tests.TestError{-32000, "rpc: can't find method \"node.getInfo\""}},
		{randomString, tests.TestError{-32000, "rpc: service/method request ill-formed: \"" + randomString + "\""}},
		{"1111", tests.TestError{-32000, "rpc: service/method request ill-formed: \"1111\""}},
		{"node getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node getSeed\""}},
		{"node^getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node^getSeed\""}},
		{"node*getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node*getSeed\""}},
		{"node&getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node&getSeed\""}},
		{"node%getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node%getSeed\""}},
		{"getSeed", tests.TestError{-32000, "rpc: service/method request ill-formed: \"getSeed\""}},
	}
	for _, tc := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  tc.Input,
		}
		getSeedWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetSeedWithBadJsonVersion(t *testing.T) {
	randomString := testutils.RandomString()
	data := []tests.Cases{
		{"1.0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{" ", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{randomString, tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"-1", tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{"3.0", tests.TestError{-32600, "jsonrpc must be 2.0"}},
	}
	for _, tc := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: tc.Input,
			Id:      apihelper.GetRequestId(),
			Method:  apihelper.GetSeedMethod,
		}
		getSeedWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetSeedWithBadRequestId(t *testing.T) {
	data := []int32{0, -1, -2147483648, 2147483647}

	for _, v := range data {
		r := insolar_api.NodeGetSeedRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      v,
			Method:  apihelper.GetSeedMethod,
		}
		getSeedRequest(t, r)
	}
}

func TestGetSeedWithTwoRequestId(t *testing.T) {
	r := insolar_api.NodeGetSeedRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      1,
		Method:  apihelper.GetSeedMethod,
	}
	getSeedRequest(t, r)
	getSeedRequest(t, r)
}
func getSeedWithBadRequest(t *testing.T, r insolar_api.NodeGetSeedRequest, error tests.TestError) {
	response, http := loggingGetSeedRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.Equal(t, error.Message, response.Error.Message)
	require.Equal(t, int32(error.Code), response.Error.Code)
	require.Empty(t, response.Result)
}
func getSeedRequest(t *testing.T, r insolar_api.NodeGetSeedRequest) string {
	response, http := loggingGetSeedRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.NotEmpty(t, response.Result)
	require.Empty(t, response.Error)
	return response.Result.Seed
}

func loggingGetSeedRequest(t *testing.T, r insolar_api.NodeGetSeedRequest) (insolar_api.NodeGetSeedResponse, *http.Response) {
	apilogger.LogApiRequest(r.Method, r, nil)
	response, http, err := apihelper.GetClient().InformationApi.GetSeed(nil, r)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	return response, http
}
