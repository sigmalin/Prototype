package sredis

import (
	"config"
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisSession struct {
	sid    string
	client *redis.Client
}

func (rs *redisSession) Set(ctx context.Context, value interface{}) error {
	return rs.client.Set(ctx, rs.sid, value, config.SESSION_DURATION).Err()
}

func (rs *redisSession) Get(ctx context.Context) interface{} {
	val, err := rs.client.Get(ctx, rs.sid).Result()
	if err != nil {
		log.Print(err)
	}
	return val
}

func (rs *redisSession) Delete(ctx context.Context) error {
	key := rs.sid
	rs.sid = ""
	return rs.client.Del(ctx, key).Err()
}

func (rs *redisSession) SessionID() string {
	return rs.sid
}

func (rs *redisSession) Update(ctx context.Context) error {
	count, err1 := rs.client.Exists(ctx, rs.sid).Result()
	if err1 != nil {
		return err1
	}

	if count == 0 {
		return fmt.Errorf("session not exist")
	}

	_, err2 := rs.client.Expire(ctx, rs.sid, config.SESSION_DURATION).Result()
	return err2
}
