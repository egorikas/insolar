package tests

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar-api/apiclient"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestGetSeed(t *testing.T) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	body := apiclient.NodeGetSeedRequest{
		Jsonrpc: "2.0",
		Id:      "2",
		Method:  "node.getSeed",
		Params:  nil,
	}
	var response, _, err = ApiClient.InformationApi.GetSeed(nil, body)
	log.Println(response)
	require.NotEqual(t, "", response.Result.Seed)
	require.Equal(t, nil, err)
}

func TestGetInfo(t *testing.T) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	body := apiclient.NetworkGetInfoRequest{
		Jsonrpc: "2.0",
		Id:      "1",
		Method:  "network.getInfo",
		Params:  nil,
	}
	var response, _, err = ApiClient.InformationApi.GetInfo(nil, body)

	log.Println(response)

	require.NotEqual(t, "", response.Result.RootDomain)
	require.NotEqual(t, "", response.Result.RootMember)
	require.NotEqual(t, "", response.Result.NodeDomain)
	require.NotEqual(t, "", response.Result.TraceID)
	require.Equal(t, nil, err)
}

func TestGetStatus(t *testing.T) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	body := apiclient.NodeGetStatusRequest{
		Jsonrpc: "2.0",
		Id:      "1",
		Method:  "node.getStatus",
		Params:  nil,
	}
	var response, _, err = ApiClient.InformationApi.GetStatus(nil, body)

	log.Println(response.Result.NetworkState)
	log.Println(response.Result)

	require.NotEqual(t, "", response.Result)
	require.Equal(t, nil, err)
}

//get seed helper

func getSeed(t *testing.T) (string, error) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	body := apiclient.NodeGetSeedRequest{
		Jsonrpc: "2.0",
		Id:      "2",
		Method:  "node.getSeed",
		Params:  nil,
	}
	var response, _, err = ApiClient.InformationApi.GetSeed(nil, body)
	log.Println(response)
	require.NotEqual(t, "", response.Result.Seed)
	require.Equal(t, nil, err)
	return response.Result.Seed, err
}
