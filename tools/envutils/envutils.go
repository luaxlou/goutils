package envutils

import (
	"github.com/joho/godotenv"
	"github.com/luaxlou/goutils/tools/fileutils"
	"log"
	"os"
)

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

//当处于开发环境时，加载当前目录Env文件
func LoadEnvOnlyDev() {

	if IsDev() {
		LoadEnv()
	}
}

//加载当前目录Env文件
func LoadEnv() {
	//多目录检查主要是为了测试代码对于环境变量的加载
	checkPaths := []string{
		"./.env",
		"../.env",
		"../../.env",
		"../../../.env",
		"../../../../.env",
		"../../../../../.env",
	}

	for _, p := range checkPaths {
		if fileutils.Exists(p) {
			err := godotenv.Load(p)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}

}
