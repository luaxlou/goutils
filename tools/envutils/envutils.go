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

func GetEnv() string {
	return currEnv
}

func IsProd() bool {
	return currEnv == Prod

}

func IsTest() bool {
	return currEnv == Test

}

func IsDev() bool {

	return currEnv == Dev

}


