package apiclient

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
