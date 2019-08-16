// +build functest

package smoke

import (
	"github.com/insolar/insolar/apitests"
	"testing"
)

func TestMain(m *testing.M) {
	apitests.StartLaunchnet(m)
}
