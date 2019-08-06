package insolarapitests

import (
	"testing"
)

func TestDepositTransfer(t *testing.T) {
	//Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//
	//seed, _ := getSeed(t)
	//ms, _ := apiclientdeprecated.NewMemberSignature()
	////create member
	//
	//body := apiclient.DepositTransferRequest{
	//	Jsonrpc: "2.0",
	//	Id:      id,
	//	Method:  "contract.call",
	//	Params: apiclient.DepositTransferRequestParams{
	//		Seed:     seed,
	//		CallSite: "deposit.transfer",
	//		CallParams: apiclient.DepositTransferRequestParamsCallParams{
	//			Amount:    "",
	//			EthTxHash: "",
	//		},
	//		PublicKey: string(ms.PemPublicKey),
	//		Reference: "", //from create member -> ref
	//	},
	//}
	//var response, _, err = ApiClient.MigrationApi.DepositTransfer(nil, "", string(ms.PemPublicKey), body)
	//id++
	//log.Println(response)
	////require.NotEqual(t, "", response.Result.Seed)
	//require.Equal(t, nil, err)
}

func TestMemberMigrationCreate(t *testing.T) {
	//Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	//
	//seed, _ := getSeed(t)
	//ms, _ := apiclientdeprecated.NewMemberSignature()
	////create member
	//
	//body := apiclient.MemberMigrationCreateRequest{
	//	Jsonrpc: "2.0",
	//	Id:      id,
	//	Method:  "contract.call",
	//	Params: apiclient.MemberMigrationCreateRequestParams{
	//		Seed:       seed,
	//		CallSite:   "member.create",
	//		CallParams: nil,
	//		PublicKey:  string(ms.PemPublicKey),
	//	},
	//}
	//var response, _, err = ApiClient.MigrationApi.MemberMigrationCreate(nil, "", string(ms.PemPublicKey), body)
	//id++
	//log.Println(response)
	////require.NotEqual(t, "", response.Result.Seed)
	//require.Equal(t, nil, err)
}
