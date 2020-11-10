package validationutils

import (
	"regexp"
	"testing"

	"github.com/luaxlou/goutils/tools/logutils"
)

func TestIsMobile(t *testing.T) {
	mobile :="16600000000"

	logutils.PrintObj(IsMobile(mobile))
}

func TestIsMobile2(t *testing.T) {
	var uriRegExpNoUser = regexp.MustCompile("^([A-Za-z]+):([^\\s;]+)(.*)$")

	r := uriRegExpNoUser.FindAllStringSubmatch("sip:057128103921@192.168.253.164:32060",-1)

	logutils.PrintObj(r)
}