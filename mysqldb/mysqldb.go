package mysqldb

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var db * gorm.DB

func init() {

	dsn := os.Getenv("DB_DSN")

	if dsn != "" {

		db = New(dsn)
	}

}

func DB() * gorm.DB {
	return db
}


func Close()  {
	if db !=nil {
		db.Close()

	}
}

func New(dsn string) *gorm.DB {

	if dsn == "" {

		dsn = "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&multiStatements=true"

	}

	return initDB(dsn)

}

func initDB(dsn string) (db *gorm.DB) {

	if dsn == "" {
		panic("dsn must not empty")

	}

	db, err := gorm.Open("mysql", dsn)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
	db.DB().SetConnMaxLifetime(3 * 60 * time.Second)

	if err != nil {
		panic("mysql connect error " + err.Error())

	}

	if db.Error != nil {
		panic("database error " + db.Error.Error())

	}

	return db
}
