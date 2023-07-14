package credis

import (
	"github.com/go-redis/redis"
)

var redisClients = make(map[int]*redis.Client)

func GetRedisClient(db int) (*redis.Client, error) {
	client, ok := redisClients[db]
	if !ok {
		conn, err := connect(db)
		if err != nil {
			return nil, err
		}
		redisClients[db] = conn
		client = conn
	}
	return client, nil
}
