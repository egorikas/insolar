package tests

import (
	"github.com/insolar/insolar/apitests/apihelper"
	"log"
	"testing"
)

var logger *log.Logger

func TestGetSeed(t *testing.T) {
	apihelper.GetSeed()
}

func TestCreateMember(t *testing.T) {
	//rootMember := apihelper.GetRootMember()
	member := apihelper.CreateMember()
	logger.Printf(member.MemberResponseResult.Result.CallResult.Reference)
}
