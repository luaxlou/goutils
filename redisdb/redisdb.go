package redisdb

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

var db *redis.Client

func init() {

	host := os.Getenv("REDIS_HOST")
	rdb := os.Getenv("REDIS_DB")
	pass := os.Getenv("REDIS_PASS")

	if host != "" {

		n, _ := strconv.Atoi(rdb)

		db = New(host, pass, n)
	}

}

func DB() *redis.Client {
	return db
}


func Close()  {
	if db !=nil {
		db.Close()

	}
}

func New(host string, pass string, db int) *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
}
