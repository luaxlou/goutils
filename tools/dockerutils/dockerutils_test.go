package dockerutils

import (
	"github.com/luaxlou/goutils/tools/logutils"
	"testing"
)

func TestGetMachineIp(t *testing.T) {

	ip :=GetMachineIp()

	logutils.PrintObj(ip)
}
