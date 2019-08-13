package apihelper

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_internal_api"
	"github.com/insolar/insolar/apitests/apihelper/apilogger"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	//migration_api
	MigrationAddAddresses = "migration.addAddresses"
	MigrationGetInfo      = "migration.getInfo"
	DepositMigration      = "deposit.migration"
	MemberGetBalance      = "member.getBalance"
	DeactivateDaemon      = "migration.deactivateDaemon"
)

var internalMemberApi = getInternalClient().MemberApi
var internalMigrationApi = getInternalClient().MigrationApi
var internalObserverApi = getInternalClient().ObserverApi

func getInternalClient() *insolar_internal_api.APIClient {
	c := insolar_internal_api.Configuration{
		BasePath: url,
	}
	return insolar_internal_api.NewAPIClient(&c)
}

func AddMigrationAddresses(t *testing.T) insolar_internal_api.MigrationDeactivateDaemonResponse {
	ms, _ := NewMemberSignature()
	uuids, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	adminPub, _ := LoadAdminMemberKeys() //todo getinfo

	body := insolar_internal_api.MigrationAddAddressesRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  ApiCall,
		Params: insolar_internal_api.MigrationAddAddressesRequestParams{
			Seed:     GetSeed(t),
			CallSite: MigrationAddAddresses,
			CallParams: insolar_internal_api.MigrationAddAddressesRequestParamsCallParams{
				MigrationAddresses: []string{uuids.String()},
			},
			PublicKey: adminPub,
			Reference: "",
		},
	}
	d, s, m := sign(body, ms.PrivateKey)
	apilogger.LogApiRequest(MigrationAddAddresses, body, m)
	response, http, err := internalMigrationApi.AddMigrationAddresses(nil, d, s, body)
	require.Nil(t, err)
	CheckResponseHasNoError(t, response)
	apilogger.LogApiResponse(http, response)
	apilogger.Printf("response id: %d", response.Id)
	return response
}

func GetMigrationInfo(t *testing.T) insolar_internal_api.MigrationGetInfoResponse {
	body := insolar_internal_api.MigrationGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  MigrationGetInfo,
		Params:  nil,
	}
	apilogger.LogApiRequest(MigrationGetInfo, body, nil)
	response, http, err := internalMigrationApi.GetInfo(nil, body)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	require.NotEmpty(t, response.Result.MigrationAdminMember)
	return response
}

func MigrationDeposit(t *testing.T) insolar_internal_api.DepositMigrationResponse {
	body := insolar_internal_api.DepositMigrationRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  ApiCall,
		Params: insolar_internal_api.DepositMigrationRequestParams{
			Seed:     GetSeed(t),
			CallSite: DepositMigration,
			CallParams: insolar_internal_api.DepositMigrationRequestParamsCallParams{
				Amount:           "1000",
				EthTxHash:        "Eth_TxHash_test",
				MigrationAddress: "", //todo getinfo
			},
			PublicKey: "", //migrationDaemonMember
			Reference: "", //migrationDaemonMember
		},
	}
	apilogger.LogApiRequest(MigrationGetInfo, body, nil)
	response, http, err := internalMigrationApi.DepositMigration(nil, "", "", body) //migrationDaemonMember
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	require.NotEmpty(t, response.Result.CallResult.MemberReference)
	return response
}

func ObserverToken(t *testing.T) map[string]interface{} {
	response, http, err := internalObserverApi.TokenGetInfo(nil)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	return response
}

func GetBalance(t *testing.T, member MemberObject) insolar_internal_api.MemberGetBalanceResponse {
	body := insolar_internal_api.MemberGetBalanceRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  ApiCall,
		Params: insolar_internal_api.MemberGetBalanceRequestParams{
			Seed:     GetSeed(t),
			CallSite: MemberGetBalance,
			CallParams: insolar_internal_api.MemberGetBalanceRequestParamsCallParams{
				Reference: member.MemberReference,
			},
			PublicKey: string(member.Signature.PemPublicKey),
			Reference: member.MemberReference,
		},
	}
	d, s, m := sign(body, member.Signature.PrivateKey)
	apilogger.LogApiRequest(MigrationGetInfo, body, m)
	response, http, err := internalMemberApi.GetBalance(nil, d, s, body)
	require.Nil(t, err)
	apilogger.LogApiResponse(http, response)
	require.NotEmpty(t, response.Result.CallResult.Balance)
	return response
}

func MigrationDeactivateDaemon(t *testing.T, migrationDaemonReference string) insolar_internal_api.MigrationDeactivateDaemonResponse {

	body := insolar_internal_api.MigrationDeactivateDaemonRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  ApiCall,
		Params: insolar_internal_api.MigrationDeactivateDaemonRequestParams{
			Seed:     GetSeed(t),
			CallSite: DeactivateDaemon,
			CallParams: insolar_internal_api.MigrationDeactivateDaemonRequestParamsCallParams{
				Reference: migrationDaemonReference, //migrationdaemon
			},
			PublicKey: "", //admin
			Reference: "", //admin
		},
	}
	//d, s, m := sign(body, ms.PrivateKey)
	apilogger.LogApiRequest(MigrationAddAddresses, body, nil)
	response, http, err := internalMigrationApi.MigrationDeactivateDaemon(nil, "", "", body)
	require.Nil(t, err)
	CheckResponseHasNoError(t, response)
	apilogger.LogApiResponse(http, response)
	apilogger.Printf("response id: %d", response.Id)
	return response
}

func getMigrationAdmin(t *testing.T) string {
	return GetMigrationInfo(t).Result.MigrationAdminMember
}
