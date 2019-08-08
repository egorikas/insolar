package apihelper

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar_internal_api"
	uuid "github.com/satori/go.uuid"
	"log"
)

const (
	//migration_api
	MIGRATIONADDADDRESSES = "migration.addAddresses"
	MIGRATIONGETINFO      = "migration.getInfo"
)

var internalMemberApi = getInternalClient().MemberApi
var internalMigrationApi = getInternalClient().MigrationApi

func getInternalClient() *insolar_internal_api.APIClient {
	c := insolar_internal_api.Configuration{
		BasePath: url,
	}
	return insolar_internal_api.NewAPIClient(&c)
}

func AddMigrationAddresses() insolar_internal_api.MigrationDeactivateDaemonResponse {
	uuids, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	adminPub, adminPrivate := LoadAdminMemberKeys()

	body := insolar_internal_api.MigrationAddAddressesRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  CONTRACTCALL,
		Params: insolar_internal_api.MigrationAddAddressesRequestParams{
			Seed:     GetSeed(),
			CallSite: MIGRATIONADDADDRESSES,
			CallParams: insolar_internal_api.MigrationAddAddressesRequestParamsCallParams{
				MigrationAddresses: []string{uuids.String()},
			},
			PublicKey: adminPub,
			Reference: getMigrationAdmin(),
		},
	}
	Logger.Printf("%v request body:\n %v", MIGRATIONADDADDRESSES, body)
	response, http, err := internalMigrationApi.AddMigrationAddresses(nil, "", adminPrivate, body) //todo подпись
	Logger.Printf("%v response statusCode:\n %v", MIGRATIONADDADDRESSES, http.StatusCode)
	Logger.Printf("%v response id:\n %v", MIGRATIONADDADDRESSES, response.Id)
	if err != nil {
		log.Fatalln(err)
	}
	Logger.Printf("%v response.Result: \n %v", MIGRATIONADDADDRESSES, response.Result)
	Logger.Printf("%v response.error: \n %v", MIGRATIONADDADDRESSES, response.Error)
	return response
}

func GetMigrationInfo() insolar_internal_api.MigrationGetInfoResponse {
	request := insolar_internal_api.MigrationGetInfoRequest{
		Jsonrpc: JSONRPCVersion,
		Id:      GetRequestId(),
		Method:  MIGRATIONGETINFO,
		Params:  nil,
	}
	Logger.Printf("%v request body:\n %v", MIGRATIONGETINFO, request)
	response, i, e := internalMigrationApi.GetInfo(nil, request)
	Logger.Printf("%v response statusCode:\n %v", MIGRATIONGETINFO, i.StatusCode)
	Logger.Printf("%v response body.Id:\n %v", MIGRATIONGETINFO, response.Id)
	Logger.Printf("%v response body.Result:\n %v", MIGRATIONGETINFO, response.Result)
	Logger.Printf("%v response body.Error:\n %v", MIGRATIONGETINFO, response.Error)
	if e != nil {
		log.Fatalln(e)
	}
	return response
}

func getMigrationAdmin() string {
	return GetMigrationInfo().Result.MigrationAdminMember
}
