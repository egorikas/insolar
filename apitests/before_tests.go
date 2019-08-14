package apitests

import (
	"github.com/insolar/insolar/functest"
	"os"
	"testing"
)

func StartLaunchnet(m *testing.M) {
	os.Exit(functest.TestsMainWrapper(m))
}
