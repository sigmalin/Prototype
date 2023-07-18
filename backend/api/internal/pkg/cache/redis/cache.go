package cacheRedis

import (
	"config"
	"context"
	"fmt"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func newCache() *cache.Cache {

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"redis1": fmt.Sprintf("%s:%d", config.REDIS_CACHE_ADDRESS, config.REDIS_CACHE_PORT),
		},
	})

	return cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, config.REDIS_CACHE_DURATION),
	})
}

func Search(key string, query func() (interface{}, error)) (interface{}, error) {

	value := new(interface{})

	cacher := GetInstance()
	err := cacher.Once(&cache.Item{
		Key:   key,
		Value: value,
		Do: func(i *cache.Item) (interface{}, error) {
			value, err := query()
			return value, err
		},
	})

	return value, err
}

func Delete(ctx context.Context, key string) error {

	cacher := GetInstance()
	return cacher.Delete(ctx, key)
}
