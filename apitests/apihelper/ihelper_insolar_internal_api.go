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
)

var internalMemberApi = getInternalClient().MemberApi
var internalMigrationApi = getInternalClient().MigrationApi

func getInternalClient() *insolar_internal_api.APIClient {
	c := insolar_internal_api.Configuration{
		BasePath: url,
	}
	return insolar_internal_api.NewAPIClient(&c)
}

func AddMigrationAddresses(t *testing.T) insolar_internal_api.MigrationDeactivateDaemonResponse {
	uuids, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	adminPub, _ := LoadAdminMemberKeys()

	body := insolar_internal_api.MigrationAddAddressesRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  ContractCall,
		Params: insolar_internal_api.MigrationAddAddressesRequestParams{
			Seed:     GetSeed(t),
			CallSite: MigrationAddAddresses,
			CallParams: insolar_internal_api.MigrationAddAddressesRequestParamsCallParams{
				MigrationAddresses: []string{uuids.String()},
			},
			PublicKey: adminPub,
			Reference: getMigrationAdmin(t),
		},
	}
	apilogger.LogApiRequest(MigrationAddAddresses, body, nil)
	response, http, err := internalMigrationApi.AddMigrationAddresses(nil, "", "", body) //todo подпись
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
	return response
}

func getMigrationAdmin(t *testing.T) string {
	return GetMigrationInfo(t).Result.MigrationAdminMember
}
