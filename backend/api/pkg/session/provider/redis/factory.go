package sredis

import (
	"github.com/go-redis/redis/v8"
)

func NewProvider(client *redis.Client) *redisProvider {
	return &redisProvider{redisClient: client}
}
