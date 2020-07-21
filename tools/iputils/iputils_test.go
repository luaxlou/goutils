package iputils

import (
	"testing"

	"github.com/luaxlou/goutils/tools/logutils"
)

func TestGetWanIp(t *testing.T) {
	ip := GetWanIp()

	logutils.PrintObj(ip)
}
