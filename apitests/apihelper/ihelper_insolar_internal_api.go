package apihelper

/*
import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_internal_api"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

const (
	//migration_api
	MIGRATIONADDADDRESSES   = "migration.addAddresses"
	MIGRATIONGETINFO   = "migration.getInfo"
)

var internalMemberApi = getInternalClient().MemberApi
var internalMigrationApi = getInternalClient().MigrationApi

func getInternalClient() *insolar_internal_api.APIClient {
	c := insolar_internal_api.Configuration{
		BasePath: "http://localhost:19101",
		//Host:     "",
	}
	return insolar_internal_api.NewAPIClient(&c)
}

func AddMigrationAddresses() (insolar_internal_api.TariffSetActiveResponse, *http.Response, error){

	uuids, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	adminPub, adminPrivate := loadAdminMemberKeys()

	body := insolar_internal_api.MigrationAddAddressesRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      getRequestId(),
		Method:  APICALL,
		Params: insolar_internal_api.MigrationAddAddressesRequestParams{
			Seed:     GetSeed(),
			CallSite: MIGRATIONADDADDRESSES,
			CallParams: insolar_internal_api.MigrationAddAddressesRequestParamsCallParams{
				Index:              1,
				MigrationAddresses: []string{uuids.String()},
			},
			PublicKey: adminPub,
			Reference: getMigrationAdmin(),
		},
	}
	return internalMigrationApi.AddMigrationAddresses(nil, "",adminPrivate, body)//todo подпись
}

func GetMigrationInfo()(insolar_internal_api.MigrationGetInfoResponse, *http.Response, error){
	return internalMigrationApi.GetInfo(nil,  insolar_internal_api.MigrationGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      getRequestId(),
		Method:  MIGRATIONGETINFO,
		Params:  nil,
	})
}

func getMigrationAdmin() string {
	response, _, err := GetMigrationInfo()
	if err != nil {
		panic(err)
	}
	return response.Result.MigrationAdminMember
}
*/
