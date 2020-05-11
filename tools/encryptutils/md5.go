package encryptutils

import (
	"crypto/md5"
	"fmt"
)

func Md5Encrypt(str string) string {

	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
