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

import "crypto"

// Define constants the requests will use:
const (
	// JSON RPC protocol version:
	JSONRPCVersion = "2.0"
)

// Declare a nested structure to form requests to Insolar's API in accordance with the specification.
// The Platform uses the basic JSON RPC 2.0 request structure:
type PlatformRequest struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Method  string         `json:"method"`
	Params  PlatformParams `json:"params"`
}

// The Platform defines params of the signed request as follows:
type PlatformParams struct {
	Seed       string      `json:"seed"`
	CallSite   string      `json:"callSite"`
	CallParams interface{} `json:"callParams"`
	Reference  string      `json:"reference"`
	PublicKey  string      `json:"publicKey"`
}

type TransferParams struct {
	Amount            string `json:"amount"`
	ToMemberReference string `json:"toMemberReference"`
}

// UserConfigJSON holds info about user
type UserConfigJSON struct {
	PrivateKey       string `json:"private_key"`
	PublicKey        string `json:"public_key"`
	Caller           string `json:"caller"`
	privateKeyObject crypto.PrivateKey
}

type GetBalanceParams struct {
	Reference string `json:"reference"`
}
