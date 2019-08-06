package apihelper

import (
	"crypto/ecdsa"
	"github.com/insolar/insolar/apitests/apiclient/insolar-api/apiclient"
	"math/big"
)

type MemberSignature struct {
	PublicKey     ecdsa.PublicKey
	PrivateKey    *ecdsa.PrivateKey
	X509PublicKey []byte
	PemPublicKey  []byte
}

type MemberObject struct {
	Signature            MemberSignature
	MemberResponseResult apiclient.MemberCreateResponse
}

type SignatureHeaders struct {
	Signature string
	Digest    string
}

type ecdsaSignature struct {
	R, S *big.Int
}
