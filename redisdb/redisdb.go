package redisdb

import (
	"github.com/go-redis/redis"
	"log"
	"os"
	"strconv"
	"time"
)

var db *redis.Client


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

	log.Println("connect to redis host",host)

	return redis.NewClient(&redis.Options{
		Addr:        host,
		Password:    pass, // no password set
		DB:          db,   // use default DB
		DialTimeout: time.Minute,
		ReadTimeout: time.Minute,
		IdleTimeout                         : time.Minute,
	})
}
