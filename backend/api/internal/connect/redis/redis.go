package credis

import (
	"fmt"
	"log"

	"config"

	"github.com/go-redis/redis"
)

func connect(db int) (*redis.Client, error) {
	conn := fmt.Sprintf("%s:%d", config.REDIS_ADDRESS, config.REDIS_PORT)

	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "",
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Print(err)
	}

	return client, err
}
