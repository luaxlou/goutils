package redisdb

import (
 	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var db *redis.Client

func InitDefault(host string) {
	db = New(host, "", 0)

}

func Init(host, pass string, n int) {
	db = New(host, pass, n)

}

func DB() *redis.Client {

	if db == nil {
		host := os.Getenv("REDIS_HOST")
		rdb := os.Getenv("REDIS_DB")
		pass := os.Getenv("REDIS_PASS")

		if host != "" {

			n, _ := strconv.Atoi(rdb)

			db = New(host, pass, n)
		}
	}
	return db
}

func Close() {
	if db != nil {
		db.Close()

	}
}

func New(host string, pass string, db int) *redis.Client {

	log.Println("connect to redis host", host)

	return redis.NewClient(&redis.Options{
		Addr:        host,
		Password:    pass, // no password set
		DB:          db,   // use default DB
		DialTimeout: time.Minute,
		ReadTimeout: time.Minute,
		IdleTimeout: time.Minute,
	})
}
