package mysqldb

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

	mysqldb := mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})

	db, err := gorm.Open(mysqldb, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Warn, // Log level
			},
		),
	})

	rawDb, _ := db.DB()

	rawDb.SetMaxIdleConns(10)
	rawDb.SetMaxOpenConns(200)
	rawDb.SetConnMaxLifetime(3 * 60 * time.Second)

	if err != nil {
		panic("mysql connect error " + err.Error())

	}

	if db.Error != nil {
		panic("database error " + db.Error.Error())

	}

	return db
}

func Close() {
	if db != nil {

		rawDb, _ := db.DB()

		if rawDb != nil {
			rawDb.Close()

		}

	}
	instances.Range(func(key, value interface{}) bool {

		g := value.(*gorm.DB)

		rawDb, _ := g.DB()

		if rawDb != nil {
			rawDb.Close()
		}

		return true

	})
}
