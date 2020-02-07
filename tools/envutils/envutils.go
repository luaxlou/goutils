package envutils

import "os"

var currEnv = "dev"

const (
	Prod = "prod"
	Test = "test"
	Dev  = "dev"
)

func init() {

	if env := os.Getenv("APP_ENV"); env != "" {

		SetEnv(env)
	}

}

func SetEnv(env string) {

	currEnv = env

}

func IsProd() bool {
	return currEnv == Test

}

func IsTest() bool {
	return currEnv == Test

}

func IsDev() bool {

	return currEnv == Dev

}
