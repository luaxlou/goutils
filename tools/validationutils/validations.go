package validationutils

import (
	"regexp"
)

//Copy from https://blog.csdn.net/qq_38572383/article/details/85776570
func IsMobile(str string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|6[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)

	return rgx.MatchString(str)

}
