package credis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClients = make(map[int]*redis.Client)

func GetRedisClient(ctx context.Context, db int) (*redis.Client, error) {
	client, ok := redisClients[db]
	if !ok {
		conn, err := connect(ctx, db)
		if err != nil {
			return nil, err
		}
		redisClients[db] = conn
		client = conn
	}
	return client, nil
}
