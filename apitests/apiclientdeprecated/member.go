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
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"
)

// Where callParams is a structure that depends on a particular method.
// The member.create request has no parameters, so it's an empty structure:
type memberCreateParams struct{}

// The result of the member.create request is as follows:
type memberResponse struct {
	PlatformResponse
	Result memberResult
}
type memberResult struct {
	CallResult struct {
		Reference string `json:"reference"`
	} `json:"callResult"`
	TraceID string `json:"TraceID"`
}

type MemberObject struct {
	Signature      MemberSignature
	MemberResponse memberResponse
}

func GetRootMember() string {
	// Form a request body for getInfo:
	getInfoReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "network.getInfo",
		ID:      id,
	}
	// Increment the id for future requests:
	id++

	// Send the request:
	infoBody := sendInfoRequest(getInfoReq)

	// Unmarshal the response:
	var info infoResponse
	err := json.Unmarshal(infoBody, &info)
	if err != nil {
		log.Fatalln(err)
	}
	// Put the rootMember reference into a variable:
	return info.Result.RootMember
}

func GetNewMember(rootMember string) MemberObject {
	var err error
	ms, _ := NewMemberSignature()
	// Form a request body for member.create:
	seed := getNewSeed()
	fmt.Println("Get seed ok   " + time.Now().Format(time.RFC850))
	createMemberReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "api.call",
		ID:      id,
		Params: PlatformParams{
			Seed:       seed,
			CallSite:   "member.create",
			CallParams: memberCreateParams{},
			Reference:  string(rootMember),
			PublicKey:  string(ms.PemPublicKey)},
	}
	// Increment the id for future requests:
	id++
	newMember := sendSignedRequest(createMemberReq, ms.PrivateKey)
	fmt.Println(string(newMember))
	var member memberResponse
	err = json.Unmarshal(newMember, &member)
	if err != nil {
		log.Fatalln(err)
	}
	// Put your reference into a variable to form a transfer request next:
	//memberReference := member.Result.CallResult.Reference
	return MemberObject{
		Signature:      ms,
		MemberResponse: member,
	}
}

func (m *MemberObject) TransferMoney(to MemberObject, amount string) []byte {
	// Get a new seed and form a transfer request:
	seed := getNewSeed()
	fmt.Println("Get seed for transfer ok   " + time.Now().Format(time.RFC850))
	// Form a request body for transfer:
	transferReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "api.call",
		ID:      id,
		Params: PlatformParams{
			Seed:     seed,
			CallSite: "member.transfer",
			CallParams: TransferParams{
				Amount:            amount,
				ToMemberReference: to.MemberResponse.Result.CallResult.Reference,
			},
			Reference: string(m.MemberResponse.Result.CallResult.Reference),
			PublicKey: string(m.Signature.PemPublicKey)},
	}

	// Send the signed transfer request:
	return sendSignedRequest(transferReq, m.Signature.PrivateKey)
}

func (m *MemberObject) GetBalance() (*big.Int, error) {
	// Get a new seed and form a transfer request:
	seed := getNewSeed()
	fmt.Println("Get seed for transfer ok   " + time.Now().Format(time.RFC850))
	// Form a request body for transfer:
	transferReq := PlatformRequest{
		JSONRPC: JSONRPCVersion,
		Method:  "api.call",
		ID:      id,
		Params: PlatformParams{
			Seed:     seed,
			CallSite: "wallet.getBalance",
			CallParams: GetBalanceParams{
				m.MemberResponse.Result.CallResult.Reference,
			}, //map[string]interface{}{"reference": m.MemberResponse.Result.CallResult.Reference},
			Reference: string(m.MemberResponse.Result.CallResult.Reference),
			PublicKey: string(m.Signature.PemPublicKey)},
	}

	// Send the signed transfer request:

	var resp ContractAnswer
	res := sendSignedRequest(transferReq, m.Signature.PrivateKey)
	err := json.Unmarshal(res, &resp)
	if err != nil {
		log.Fatalln(err)
	}
	amount, ok := new(big.Int).SetString(resp.Result.ContractResult.(map[string]interface{})["balance"].(string), 10)
	if !ok {
		return nil, fmt.Errorf("can't parse input amount for get balance")
	}
	return amount, nil
}

/*func getBalance(caller *user, reference string) (*big.Int, error) {
	res, err := signedRequest(caller, "wallet.getBalance", map[string]interface{}{"reference": reference})
	if err != nil {
		return nil, err
	}
	amount, ok := new(big.Int).SetString(res.(map[string]interface{})["balance"].(string), 10)
	if !ok {
		return nil, fmt.Errorf("can't parse input amount")
	}
	return amount, nil
}*/
