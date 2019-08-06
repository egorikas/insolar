package tests

import (
	"github.com/insolar/insolar/apitests/apiclient/insolar-api/apiclient"
	"log"
)

var Logger *log.Logger
var ApiClient *apiclient.APIClient

func init() {
	var cfg = apiclient.NewConfiguration()
	cfg.BasePath = "http://localhost:19101"
	ApiClient = apiclient.NewAPIClient(cfg)
}
