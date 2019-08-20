package apitests

import (
	"os"

	"github.com/insolar/insolar/testutils/launchnet"

	"testing"
)

// var IsLaunchnetRunning bool
var resultRun []int

func RunTestsInLaunchNet(m ...*testing.M) {
	exit := launchnet.Run(func() int {
		for _, suit := range m {
			resultRun = append(resultRun, suit.Run())
		}
		return 0
	})
	// TODO handle exit codes
	os.Exit(exit)
}
