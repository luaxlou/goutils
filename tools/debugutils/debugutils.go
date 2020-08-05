package debugutils

import (
	"log"
	"os"
)

var isDebug = false

func init() {

	if os.Getenv("DEBUG") == "1" {
		isDebug = true
	}
}

func SetDebug(flag bool) {
	isDebug = flag
}

func IsDebugging() bool {

	return isDebug
}

func Println(v ...interface{}) {
	if isDebug {

		log.Println(v...)
	}
}

func Printf(format string, v ...interface{}) {

	if isDebug {
		log.Printf(format, v...)

	}

}
