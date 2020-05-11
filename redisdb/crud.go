package redisdb

import (
	"encoding/json"
	"errors"
	"time"
)

//Must  set REDIS_HOST
func SetObj(key string, o interface{}, expiration time.Duration) {

	str, _ := json.Marshal(o)

	db.Set(key, str, expiration)

}

func GetObj(key string, o interface{}) error {

	str := db.Get(key).Val()

	if str == "" {
		return errors.New("not find")
	}

	return json.Unmarshal([]byte(str), &o)

}
