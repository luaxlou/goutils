package wechatapp

import (
	"testing"

	"github.com/luaxlou/goutils/tools/logutils"
)

func TestDecryptEncryptedData(t *testing.T) {

	sessionKey := ""
	encryptedData := ""
	iv := ""

	f, err := DecryptPhoneNumber(sessionKey, encryptedData, iv)

	logutils.PrintObj(err)
	logutils.PrintObj(f)

}
