package apihelper

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/insolar/insolar/apitests/apiclient/insolar-api/apiclient"
	"log"
	"os"
)

var id int32 = 0

const (
	url            = "http://localhost:19101"
	JSONRPCVersion = "2.0"
	APICALL        = "api.call"
	GETSEED        = "node.getSeed"
	GETINFO        = "network.getInfo"
	MEMBERCREATE   = "member.create"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
var informationApi = getClient().InformationApi
var memberApi = getClient().MemberApi

func getClient() *apiclient.APIClient {
	c := apiclient.Configuration{
		BasePath: "http://localhost:19101",
		//Host:     "",
	}
	return apiclient.NewAPIClient(&c)
}

func getRequestId() int32 {
	id++
	return id
}

func GetSeed() string {
	r := apiclient.NodeGetSeedRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      getRequestId(),
		Method:  GETSEED,
	}
	response, _, err := informationApi.GetSeed(nil, r)
	if err != nil {
		logger.Fatalln(err)
	}
	s := response.Result.Seed
	logger.Println("Get seed result: " + s)
	return s
}

func GetInfo() apiclient.NetworkGetInfoResponseResult {
	infoBody := apiclient.NetworkGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      getRequestId(),
		Method:  GETINFO,
		Params:  nil,
	}
	response, _, err := informationApi.GetInfo(nil, infoBody)
	if err != nil {
		logger.Fatalln(err)
	}
	return response.Result
}

func GetRootMember() string {
	return GetInfo().RootMember
}

func CreateMember() MemberObject {
	var err error
	ms, _ := NewMemberSignature()
	seed := GetSeed()

	request := apiclient.MemberCreateRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      1,
		Method:  APICALL,
		Params: apiclient.MemberCreateRequestParams{
			Seed:       seed,
			CallSite:   MEMBERCREATE,
			CallParams: nil,
			PublicKey:  string(ms.PemPublicKey),
		},
	}
	//json.Marshal(request)

	var headers = SignRequestHeaders(request, ms.PrivateKey)
	response, _, err := memberApi.MemberCreate(nil, headers.Digest, headers.Signature, request)
	if err != nil {
		log.Fatalln(err)
	}
	// Put your reference into a variable to form a transfer request next:
	//memberReference := member.Result.CallResult.Reference
	return MemberObject{
		Signature:            ms,
		MemberResponseResult: response,
	}
}

func SignRequestHeaders(payload interface{}, privateKey *ecdsa.PrivateKey) SignatureHeaders {
	// Marshal the payload into JSON:
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}

	// Take a SHA-256 hash of the payload:
	hash := sha256.Sum256(jsonPayload)

	// Sign the hash with the private key:
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Fatalln(err)
	}

	// See if the signature is valid:
	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("signature verified:", valid)

	// Convert the signature into ASN.1 format:
	sig := ecdsaSignature{
		R: r,
		S: s,
	}
	signature, _ := asn1.Marshal(sig)

	// Convert both hash and signature into a Base64 string:
	hash64 := base64.StdEncoding.EncodeToString(hash[:])
	signature64 := base64.StdEncoding.EncodeToString(signature)

	// Set headers and send the signed request:
	var Digest = "SHA-256=" + hash64
	var Signature = "keyId=\"member-pub-key\", algorithm=\"ecdsa\", headers=\"digest\", signature=" + signature64
	fmt.Printf("Digest %v\n", Digest)
	fmt.Printf("Signature %v\n", Signature)
	return SignatureHeaders{
		Signature: Signature,
		Digest:    Digest,
	}
}
