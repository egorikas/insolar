package apiclient

import (
	"context"
	"fmt"
	"github.com/insolar/insolar/apitests/apiclient/insolar-api/openapi"
)

func main() {
	//var member_api openapi.MemberApiService
	var infoApi openapi.InformationApiService
	//var migration_api openapi.MigrationApiService

	var req = openapi.NetworkGetInfoRequest{
		Jsonrpc: "",
		Id:      nil,
		Method:  "",
		Params:  nil,
	}

	fmt.Printf(req.Jsonrpc)
	infoApi.GetInfo(context.Background(), req)

}
