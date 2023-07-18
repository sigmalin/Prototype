package credis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func connect(ctx context.Context, addr string) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Print(err)
	}

	return client, err
}
