package ip2location

import (
	"log"
	"testing"

	"github.com/luaxlou/goutils/tools/logutils"
)

func TestGetLocation(t *testing.T) {

	info, err := GetLocation("180.97.33.136")

	if err != nil {

		log.Println(err)
	}

	logutils.PrintObj(info)
}
