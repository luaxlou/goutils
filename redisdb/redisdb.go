package redisdb

import "github.com/go-redis/redis"

func New(host string, pass string, db string) *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pass, // no password set
		DB:       db,   // use default DB
	})
}
