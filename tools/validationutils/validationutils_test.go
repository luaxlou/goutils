package validationutils

import (
	"testing"

	"github.com/luaxlou/goutils/tools/logutils"
)

func TestIsMobile(t *testing.T) {
	mobile :="16600000000"

	logutils.PrintObj(IsMobile(mobile))
}
