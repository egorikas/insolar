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

func TestGetInfo(t *testing.T) {
	response := apihelper.GetInfo(t)
	require.NotEmpty(t, response.RootDomain)
	require.NotEmpty(t, response.RootMember)
	require.NotEmpty(t, response.NodeDomain)
	require.NotEmpty(t, response.TraceID)
}

func TestGetInfoWithBadMethod(t *testing.T) {
	randomString := testutils.RandomString()
	data := []tests.Cases{
		{"", tests.TestError{-32000, "rpc: service/method request ill-formed: \"\""}},
		{" ", tests.TestError{-32000, "rpc: service/method request ill-formed: \" \""}},
		{"node.getInfo", tests.TestError{-32000, "rpc: can't find method \"node.getInfo\""}},
		{randomString, tests.TestError{-32000, "rpc: service/method request ill-formed: \"" + randomString + "\""}},
		{"1111", tests.TestError{-32000, "rpc: service/method request ill-formed: \"1111\""}},
		{"node getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node getInfo\""}},
		{"node^getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node^getInfo\""}},
		{"node*getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node*getInfo\""}},
		{"node&getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node&getInfo\""}},
		{"node%getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"node%getInfo\""}},
		{"getInfo", tests.TestError{-32000, "rpc: service/method request ill-formed: \"getInfo\""}},
	}
	for _, tc := range data {
		r := insolar_api.NetworkGetInfoRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      apihelper.GetRequestId(),
			Method:  tc.Input,
		}
		getInfoWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetInfoWithoutMethod(t *testing.T) {
	r := insolar_api.NetworkGetInfoRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      apihelper.GetRequestId(),
	}
	getInfoWithBadRequest(t, r, tests.TestError{-32000, "rpc: service/method request ill-formed: \"\""}) //todo bug - need error: "method name is required"
}

func TestGetInfoWithBadJsonVersion(t *testing.T) {
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
		r := insolar_api.NetworkGetInfoRequest{
			Jsonrpc: tc.Input,
			Id:      apihelper.GetRequestId(),
			Method:  apihelper.GetInfoMethod,
		}
		getInfoWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetInfoWithoutJsonField(t *testing.T) {
	r := insolar_api.NetworkGetInfoRequest{
		Id:     apihelper.GetRequestId(),
		Method: apihelper.GetInfoMethod,
	}
	getInfoWithBadRequest(t, r, tests.TestError{-32600, "jsonrpc must be 2.0"})
}

func TestGetInfoWithBadRequestId(t *testing.T) {
	data := []tests.CasesInt{
		//{0, error{-32600,"jsonrpc must be 2.0"},},//todo какие требования к id?
		//{-1, error{-32600,"jsonrpc must be 2.0"},},//todo какие требования к id?
		{-2147483648, tests.TestError{-32600, "jsonrpc must be 2.0"}},
		{2147483647, tests.TestError{-32600, "jsonrpc must be 2.0"}},
	}
	for _, tc := range data {
		r := insolar_api.NetworkGetInfoRequest{
			Jsonrpc: apihelper.JSONRPCVersion,
			Id:      tc.Input,
			Method:  apihelper.GetInfoMethod,
		}
		getInfoWithBadRequest(t, r, tc.ExpectedError)
	}
}

func TestGetInfoWithoutRequestId(t *testing.T) { //по умолчанию id = 0 //todo
	r := insolar_api.NetworkGetInfoRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Method:  apihelper.GetInfoMethod,
	}
	getInfoWithBadRequest(t, r, tests.TestError{-32600, "jsonrpc must be 2.0"})
}

func TestGetInfoWithParams(t *testing.T) {
	//type any interface{}
	//var args map[string]any
	//r := insolar_api.NetworkGetInfoRequest{
	//	Jsonrpc: apihelper.JSONRPCVersion,
	//	Id:      apihelper.GetRequestId(),
	//	Method:  apihelper.GetInfoMethod,
	//	Params:  args,
	//}
	//getInfoWithBadRequest(t, r, error{-32600, "jsonrpc must be 2.0"})
}

func TestGetInfoWithTwoRequestId(t *testing.T) {
	data := []tests.CasesInt{
		{1, tests.TestError{}},
		{1, tests.TestError{-32600, "jsonrpc must be 2.0"}},
	}
	r := insolar_api.NetworkGetInfoRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      data[0].Input,
		Method:  apihelper.GetInfoMethod,
	}
	getInfoRequest(t, r)
	r2 := insolar_api.NetworkGetInfoRequest{
		Jsonrpc: apihelper.JSONRPCVersion,
		Id:      data[1].Input,
		Method:  apihelper.GetInfoMethod,
	}
	getInfoWithBadRequest(t, r2, data[1].ExpectedError) //todo одинаковые id это нормально?

}
func getInfoWithBadRequest(t *testing.T, r insolar_api.NetworkGetInfoRequest, error tests.TestError) {
	response, http := loggingGetInfoRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.Equal(t, error.Message, response.Error.Message)
	require.Equal(t, int32(error.Code), response.Error.Code)
	require.Empty(t, response.Result)
}
func getInfoRequest(t *testing.T, r insolar_api.NetworkGetInfoRequest) insolar_api.NetworkGetInfoResponseResult {
	response, http := loggingGetInfoRequest(t, r)
	require.Equal(t, 200, http.StatusCode)
	require.NotEmpty(t, response.Result.TraceID)
	require.NotEmpty(t, response.Result.RootMember)
	require.NotEmpty(t, response.Result.RootDomain)
	require.NotEmpty(t, response.Result.NodeDomain)
	require.Empty(t, response.Error)
	return response.Result
}

func loggingGetInfoRequest(t *testing.T, r insolar_api.NetworkGetInfoRequest) (insolar_api.NetworkGetInfoResponse, *http.Response) {
	apilogger.LogApiRequest(r.Method, r, nil)
	response, http, err := apihelper.GetClient().InformationApi.GetInfo(nil, r)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	return response, http
}
