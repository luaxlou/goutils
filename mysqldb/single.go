package mysqldb

import (
	"os"

	"gorm.io/gorm"
)

//使用单例模式======

var db *gorm.DB

func Init(dsn string) {
	db = New(dsn)

}

func DB() *gorm.DB {

	if db == nil {
		dsn := os.Getenv("DB_DSN")

		if dsn != "" {

			db = New(dsn)
		}
	}

	return db
}

