package apihelper

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/insolar/insolar/apitests/apiclient/insolar_api"
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

type ecdsaSignature struct {
	R, S *big.Int
}
