package cacheRedis

import (
	"github.com/go-redis/cache/v8"
)

var instance *cache.Cache

func GetInstance() *cache.Cache {
	if instance == nil {
		instance = newCache()
	}
	return instance
}
