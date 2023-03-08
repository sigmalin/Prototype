package redis

import (
	"fmt"
	"log"

	"custom/config"

	goRedis "github.com/go-redis/redis"
)

var redisClients = make(map[int]*goRedis.Client)

func connect(db int) (*goRedis.Client, error) {
	conn := fmt.Sprintf("%s:%d", config.REDIS_ADDRESS, config.REDIS_PORT)

	client := goRedis.NewClient(&goRedis.Options{
		Addr:     conn,
		Password: "",
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func GetRedisClient(db int) (*goRedis.Client, error) {
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
