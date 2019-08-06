package tests

import (
	"fmt"
	"github.com/insolar/insolar/apitests/apihelper"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var logger *log.Logger

func TestGetSeed(t *testing.T) {
	seed := apihelper.GetSeed()
	fmt.Printf(seed)
	require.NotEmpty(t, seed)
}

func TestCreateMember(t *testing.T) {
	//rootMember := apihelper.GetRootMember()
	member := apihelper.CreateMember()
	logger.Printf(member.MemberResponseResult.Result.CallResult.Reference)
}
