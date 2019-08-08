package apihelper

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

var id int32 = 0

const (
	url            = "http://localhost:19101"
	JSONRPCVersion = "2.0"
	APICALL        = "api.call"
	CONTRACTCALL   = "contract.call"
	//information_api
	GETSEED   = "node.getSeed"
	GETINFO   = "network.getInfo"
	GETSTATUS = "node.getStatus"
	//member_api
	MEMBERCREATE   = "member.create"
	MEMBERTRANSFER = "member.transfer"
	//migration_api
	MEMBERMIGRATIONCREATE = "member.migrationCreate"
	DEPOSITTRANSFER       = "deposit.transfer"
)

var Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var informationApi = GetClient().InformationApi
var memberApi = GetClient().MemberApi
var migrationApi = GetClient().MigrationApi

func GetClient() *insolar_api.APIClient {
	c := insolar_api.Configuration{
		BasePath: url,
	}
	return insolar_api.NewAPIClient(&c)
}

func GetRequestId() int32 {
	id++
	return id
}

func GetSeed(t *testing.T) string {
	r := insolar_api.NodeGetSeedRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  GETSEED,
	}
	return GetSeedRequest(t, r)
}

func GetSeedRequest(t *testing.T, r insolar_api.NodeGetSeedRequest) string {
	Logger.Printf("%v request body:\n %v", GETSEED, r)
	response, http, err := informationApi.GetSeed(nil, r)
	checkResponseHasNoError(t, response)
	require.Nil(t, err)
	Logger.Printf("%v response statusCode:\n %v", GETSEED, http.StatusCode)
	Logger.Printf("%v response id:\n %v", GETSEED, response.Id)
	Logger.Printf("%v response body:\n %v", GETSEED, response)
	Logger.Printf("%v response Err:\n %v", GETSEED, response.Error)
	return response.Result.Seed
}

func GetInfo(t *testing.T) insolar_api.NetworkGetInfoResponseResult {
	infoBody := insolar_api.NetworkGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  GETINFO,
		Params:  nil,
	}
	response, _, err := informationApi.GetInfo(nil, infoBody)
	if err != nil {
		Logger.Fatalln(err)
	}
	Logger.Println("Get info result: ok")
	checkResponseHasNoError(t, response)

	return response.Result
}

func GetStatus(t *testing.T) insolar_api.NodeGetStatusResponseResult {
	infoBody := insolar_api.NodeGetStatusRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  GETSTATUS,
		Params:  nil,
	}
	response, _, err := informationApi.GetStatus(nil, infoBody)
	require.Nil(t, err)
	checkResponseHasNoError(t, response)

	return response.Result
}

func GetRootMember(t *testing.T) string {
	return GetInfo(t).RootMember
}

func CreateMember(t *testing.T) MemberObject {
	var err error
	ms, _ := NewMemberSignature()
	seed := GetSeed(t)

	request := insolar_api.MemberCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  APICALL,
		Params: insolar_api.MemberCreateRequestParams{
			Seed:      seed,
			CallSite:  MEMBERCREATE,
			PublicKey: string(ms.PemPublicKey),
		},
	}
	d, s := sign(request, ms.PrivateKey)
	response, _, err := memberApi.MemberCreate(nil, d, s, request)
	require.Nil(t, err)
	checkResponseHasNoError(t, response)
	Logger.Printf("Member created: %v", response.Result.CallResult.Reference)
	return MemberObject{
		MemberReference:      response.Result.CallResult.Reference,
		Signature:            ms,
		MemberResponseResult: response,
	}
}

func (m *MemberObject) Transfer(t *testing.T, toMemberRef string, amount string) insolar_api.MemberTransferResponse {
	seed := GetSeed(t)
	request := insolar_api.MemberTransferRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  APICALL,
		Params: insolar_api.MemberTransferRequestParams{
			Seed:     seed,
			CallSite: MEMBERTRANSFER,
			CallParams: insolar_api.MemberTransferRequestParamsCallParams{
				Amount:            amount,
				ToMemberReference: toMemberRef,
			},
			PublicKey: string(m.Signature.PemPublicKey),
			Reference: m.MemberResponseResult.Result.CallResult.Reference,
		},
	}
	d, s := sign(request, m.Signature.PrivateKey)
	response, _, err := memberApi.MemberTransfer(nil, d, s, request)
	require.Nil(t, err)
	checkResponseHasNoError(t, response)
	return response
}

func MemberMigrationCreate(t *testing.T) MemberObject {
	var err error
	ms, _ := NewMemberSignature()
	seed := GetSeed(t)

	request := insolar_api.MemberMigrationCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      1,
		Method:  APICALL,
		Params: insolar_api.MemberMigrationCreateRequestParams{
			Seed:       seed,
			CallSite:   MEMBERMIGRATIONCREATE,
			CallParams: nil,
			PublicKey:  string(ms.PemPublicKey),
		},
	}
	//json.Marshal(request)

	d, s := sign(request, ms.PrivateKey)
	Logger.Printf("%v request body:\n %v", MEMBERMIGRATIONCREATE, request)
	response, http, err := migrationApi.MemberMigrationCreate(nil, d, s, request)
	Logger.Printf("%v response body:\n %v", MEMBERMIGRATIONCREATE, response)
	Logger.Printf("%v response Status:\n %v", MEMBERMIGRATIONCREATE, http.StatusCode)
	checkResponseHasNoError(t, response)
	if err != nil {
		log.Fatalln(err)
	}
	// Put your reference into a variable to form a transfer request next:
	//memberReference := member.Result.CallResult.Reference

	return MemberObject{
		MemberReference: response.Result.CallResult.Reference,
		Signature:       ms,
		MemberResponseResult: insolar_api.MemberCreateResponse{
			Jsonrpc: response.Jsonrpc,
			Id:      response.Id,
			Result: insolar_api.MemberCreateResponseResult{
				CallResult: insolar_api.MemberCreateResponseResultCallResult{
					Reference: response.Result.CallResult.Reference,
				},
				RequestReference: response.Result.RequestReference,
				TraceID:          response.Result.TraceID,
			},
			Error: insolar_api.MemberCreateResponseError{
				Data: insolar_api.MemberCreateResponseErrorData{
					RequestReference: response.Error.Data.RequestReference,
					TraceID:          response.Error.Data.TraceID,
				},
				Code:    response.Error.Code,
				Message: response.Error.Message,
			},
		},
	}
}

func DepositTransfer(t *testing.T) insolar_api.DepositTransferResponse {
	var err error
	ms, _ := NewMemberSignature()

	request := insolar_api.DepositTransferRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      1,
		Method:  APICALL,
		Params: insolar_api.DepositTransferRequestParams{
			Seed:     GetSeed(t),
			CallSite: DEPOSITTRANSFER,
			CallParams: insolar_api.DepositTransferRequestParamsCallParams{
				Amount:    "1000",
				EthTxHash: "",
			},
			PublicKey: string(ms.PemPublicKey),
		},
	}

	d, s := sign(request, ms.PrivateKey)
	Logger.Printf("%v request body:\n %v", DEPOSITTRANSFER, request)
	response, http, err := migrationApi.DepositTransfer(nil, d, s, request)
	Logger.Printf("%v response body:\n %v", DEPOSITTRANSFER, response)
	Logger.Printf("%v response Status:\n %v", DEPOSITTRANSFER, http.StatusCode)
	checkResponseHasNoError(t, response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}
