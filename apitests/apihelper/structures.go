package apihelper

import (
	"crypto/ecdsa"
	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
	"math/big"
)

type MemberSignature struct {
	PublicKey     ecdsa.PublicKey
	PrivateKey    *ecdsa.PrivateKey
	X509PublicKey []byte
	PemPublicKey  []byte
}

type MemberObject struct {
	MemberReference      string
	Signature            MemberSignature
	MemberResponseResult insolar_api.MemberCreateResponse200
}

type SignatureHeaders struct {
	Signature string
	Digest    string
}

type ecdsaSignature struct {
	R, S *big.Int
}
