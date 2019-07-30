package tests

import (
	"github.com/insolar/insolar/apitests/apiclient"
	"github.com/insolar/insolar/apitests/introspection"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"fmt"
	// - Big numbers to store signatures.
)

var Wg sync.WaitGroup
var Logger *log.Logger

// The identifier is to be incremented in every request and each response will contain a corresponding one.
// The transfer request sends an amount of funds to member identified by a reference:
func TestCreateTransferGetBalance(t *testing.T) {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	rootMember := apiclient.GetRootMember()
	member1 := apiclient.GetNewMember(rootMember)
	member2 := apiclient.GetNewMember(rootMember)

	count := 5
	Wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go transfer(member1, member2)
		go transfer(member2, member1)
	}
	Wg.Wait()
	Logger.Println("finished")
}

func transfer(memberFrom apiclient.MemberObject, memberTo apiclient.MemberObject) {
	defer Wg.Done()
	result := memberFrom.TransferMoney(memberTo, "100")
	Logger.Printf("Transfer from: %v to %v",
		memberFrom.MemberResponse.Result.CallResult.Reference,
		memberTo.MemberResponse.Result.CallResult.Reference)
	Logger.Println("result: " + string(result))
	time.Sleep(100 + time.Millisecond)
}

func TestJetSplit(t *testing.T) {
	var jetCount, tmpInt int

	jetCount = 0
	for _, mycount := range introspection.GetMessageNodeCounters("127.0.0.1:55503").Counters {
		if mycount.Name == "TypeHotObjects" {
			tmpInt, _ = strconv.Atoi(mycount.Count)
			jetCount = (jetCount) + tmpInt
		}
	}

	for _, mycount2 := range introspection.GetMessageNodeCounters("127.0.0.1:55505").Counters {
		if mycount2.Name == "TypeHotObjects" {
			tmpInt, _ = strconv.Atoi(mycount2.Count)
			jetCount = (jetCount) + tmpInt
		}
	}

	var nodeStatus1 apiclient.NodeStatus
	nodeStatus1 = apiclient.GetNodeStatus("http://localhost:19103/api/")
	nodeStatus1 = apiclient.GetNodeStatus("http://localhost:19105/api/")
	fmt.Println(nodeStatus1.Result.PulseNumber)
	fmt.Printf("%s%d%s", "Jets Count: ", jetCount, "\n")
}
