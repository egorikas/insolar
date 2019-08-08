package apihelper

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"github.com/insolar/insolar/api"
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var id int32 = 0

const (
	url            = "http://localhost:19101"
	JSONRPCVersion = "2.0"
	APICALL        = "api.call"
	//information_api
	GETSEED   = "node.getSeed"
	GETINFO   = "network.getInfo"
	GETSTATUS = "node.getStatus"
	//member_api
	MEMBERCREATE   = "member.create"
	MEMBERTRANSFER = "member.transfer"
	//migration_api
	MEMBERMIGRATIONCREATE = "member.migrationCreate"
)

var Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var informationApi = GetClient().InformationApi
var memberApi = GetClient().MemberApi
var migrationApi = GetClient().MigrationApi

func GetClient() *insolar_api.APIClient {
	c := insolar_api.Configuration{
		BasePath: "http://localhost:19101",
		//Host:     "",
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
	response, _, err := informationApi.GetSeed(nil, r)
	if err != nil {
		Logger.Fatalln(err)
	}
	s := response.Result.Seed
	Logger.Println("Get seed result: " + s)
	return s
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
	checkResponseHasNoError(t, response)
	require.Nil(t, err)
	Logger.Printf("Member created: %v", response.Result.CallResult.Reference)
	return MemberObject{
		Signature:            ms,
		MemberResponseResult: response,
	}
}

func (m *MemberObject) Transfer(t *testing.T, toMemberRef string, amount string) string {
	//seed := GetSeed()
	//request := insolar_api.MemberTransferRequest{
	//	Jsonrpc: JSONRPCVersion,
	//	Id:      getRequestId(),
	//	Method:  APICALL,
	//	Params: insolar_api.MemberTransferRequestParams{
	//		Seed:     seed,
	//		CallSite: MEMBERTRANSFER,
	//		CallParams: insolar_api.MemberTransferRequestParamsCallParams{
	//			Amount:            amount,
	//			ToMemberReference: toMemberRef,
	//		},
	//		PublicKey: string(m.Signature.PemPublicKey),
	//		Reference: m.MemberResponseResult.Result.CallResult.Reference,
	//	},
	//}
	//d, s := sign(request, m.Signature.PrivateKey)
	//response, _, err := memberApi.MemberTransfer(nil, d, s, request)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	return "nil"
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
	response, _, err := migrationApi.MemberMigrationCreate(nil, d, s, request)
	checkResponseHasNoError(t, response)
	if err != nil {
		log.Fatalln(err)
	}
	// Put your reference into a variable to form a transfer request next:
	//memberReference := member.Result.CallResult.Reference

	return MemberObject{
		Signature: ms,
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

func sign(payload interface{}, privateKey *ecdsa.PrivateKey) (string, string) {
	var err error
	// get hash of byte slice of the payload encoded with the same way as openapi-generator does in the generated client.
	// this is done to avoid setting incorrect body value into request by generated code.
	// if you use custom code to create insolar-api client, use 'json.Marshal(payload)' and get hash value of it s result.
	bodyBuf := &bytes.Buffer{}
	err = json.NewEncoder(bodyBuf).Encode(payload)
	if err != nil {
		log.Fatalln(err)
	}
	request, err := http.NewRequest("ignore", "ignore", bodyBuf)
	memberCreateRequest := reflect.TypeOf(payload)
	rawBody, err := api.UnmarshalRequest(request, &memberCreateRequest)
	if err != nil {
		Logger.Fatalln(err)
	}
	hash := sha256.Sum256(rawBody)

	// Sign the hash with the private key:
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		Logger.Fatalln(err)
	}

	// See if the signature is valid:
	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	if !valid {
		Logger.Fatal("signature not verified")
	}

	// Convert the signature into ASN.1 format:
	sig := ecdsaSignature{
		R: r,
		S: s,
	}
	signature, _ := asn1.Marshal(sig)

	// Convert both hash and signature into a Base64 string:
	hash64 := base64.StdEncoding.EncodeToString(hash[:])
	signature64 := base64.StdEncoding.EncodeToString(signature)

	var Digest = "SHA-256=" + hash64
	var Signature = "keyId=\"member-pub-key\", algorithm=\"ecdsa\", headers=\"digest\", signature=" + signature64
	Logger.Println("Digest = " + Digest)
	Logger.Println("Signature = " + Signature)
	return Digest, Signature
}
