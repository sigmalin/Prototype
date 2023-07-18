package credis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClients = make(map[string]*redis.Client)

func GetRedisClient(ctx context.Context, addr string) (*redis.Client, error) {
	client, ok := redisClients[addr]
	if !ok {
		conn, err := connect(ctx, addr)
		if err != nil {
			return nil, err
		}
		redisClients[addr] = conn
		client = conn
	}
	return client, nil
}
