//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package apiclientdeprecated

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

// Create a variable for the JSON RPC 2.0 request identifier:
var id int = 1

// The identifier is to be incremented in every request and each response will contain a corresponding one.

const (
	// Endpoint URL for local deployment (is to be changed to a production URL):
	url = "http://localhost:19101/api/"
)

// Declare a structure to contain the ECDSA signature:
type ecdsaSignature struct {
	R, S *big.Int
}

// Create and initialize an HTTP client for connection re-use:
var ApiHttpClient *http.Client

func init() {
	ApiHttpClient = &http.Client{}
}

type MemberSignature struct {
	PublicKey     ecdsa.PublicKey
	PrivateKey    *ecdsa.PrivateKey
	X509PublicKey []byte
	PemPublicKey  []byte
}

func NewMemberSignature() (MemberSignature, error) {
	var err error
	privateKey := new(ecdsa.PrivateKey)
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return MemberSignature{}, err
	}
	var publicKey ecdsa.PublicKey
	publicKey = privateKey.PublicKey
	// Convert the public key into PEM format:
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return MemberSignature{}, err
	}
	pemPublicKey := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509PublicKey})

	return MemberSignature{
		PublicKey:     publicKey,
		PrivateKey:    privateKey,
		X509PublicKey: x509PublicKey,
		PemPublicKey:  pemPublicKey,
	}, err
}

func SendPlatformRequest(payload PlatformRequest, intUrl string, method string, contentType string) []byte {
	// Marshal the payload into JSON:
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	return SendRequest(jsonPayload, intUrl, method, contentType)
}

// Create a function to send information requests:
func SendRequest(jsonPayload []byte, intUrl string, method string, contentType string) []byte {
	// Create a new HTTP request and send it:
	request, err := http.NewRequest(method, intUrl+"rpc", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("ContentType", contentType)
	response, err := ApiHttpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer request.Body.Close()

	// Receive and return the response body:
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseBody))
	id++
	return responseBody
}

// Create a function to send signed requests:
func sendSignedRequest(payload PlatformRequest, privateKey *ecdsa.PrivateKey) []byte {
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
	request, err := http.NewRequest("POST", url+"call", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("ContentType", "application/json")

	// Put the hash string into the HTTP Digest header:
	request.Header.Set("Digest", "SHA-256="+hash64)

	// Put the signature string into the HTTP Signature header:
	request.Header.Set("Signature", "keyId=\"member-pub-key\", algorithm=\"ecdsa\", headers=\"digest\", signature="+signature64)
	fmt.Println(request.Header)
	fmt.Println(request.Body)

	// Send the signed request:
	response, err := ApiHttpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	// Receive the response body:
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	id++
	return responseBody
}

// Create a function to get a new seed for each signed request:
func getNewSeed() string {
	// Form a request body for getSeed:
	getSeedReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "node.getSeed",
		ID:      id,
	}
	// Increment the id for future requests:
	id++

	// Create a variable to unmarshal the seed response into:
	var seed seedResponse
	// Send the seed request:
	seedRespBody := sendInfoRequest(getSeedReq)
	// Unmarshal the response:
	err := json.Unmarshal(seedRespBody, &seed)
	if err != nil {
		log.Fatalln(err)
	}
	// Put the current seed into a variable:
	return seed.Result.Seed
}

// Create a function to send information requests:
func sendInfoRequest(payload PlatformRequest) []byte {
	return SendPlatformRequest(payload, url, "POST", "application/json")
}

func sendUrlRequest(url string) []byte {

	// Create a new HTTP request and send it:
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatalln(err)
	}
	response, err := ApiHttpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer request.Body.Close()

	// Receive and return the response body:
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseBody))
	return responseBody
}
