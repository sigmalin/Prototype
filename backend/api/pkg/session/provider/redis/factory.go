package sredis

import (
	"github.com/go-redis/redis"
)

func NewProvider(client *redis.Client) *redisProvider {
	return &redisProvider{redisClient: client}
}
