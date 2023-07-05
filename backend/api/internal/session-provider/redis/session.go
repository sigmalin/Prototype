package sredis

import (
	"config"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

type redisSession struct {
	sid    string
	client *redis.Client
}

func (rs *redisSession) Set(value interface{}) error {
	return rs.client.Set(rs.sid, value, config.SESSION_DURATION).Err()
}

func (rs *redisSession) Get() interface{} {
	val, err := rs.client.Get(rs.sid).Result()
	if err != nil {
		log.Print(err)
	}
	return val
}

func (rs *redisSession) Delete() error {
	key := rs.sid
	rs.sid = ""
	return rs.client.Del(key).Err()
}

func (rs *redisSession) SessionID() string {
	return rs.sid
}

func (rs *redisSession) Update() error {
	count, err1 := rs.client.Exists(rs.sid).Result()
	if err1 != nil {
		return err1
	}

	if count == 0 {
		return fmt.Errorf("session not exist")
	}

	_, err2 := rs.client.Expire(rs.sid, config.SESSION_DURATION).Result()
	return err2
}
