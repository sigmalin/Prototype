package initial

import (
	"context"
	"log"
	"sync"
	"time"

	connector "connect/redis"
)

func initRedis(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start redis")

	ctx := context.Background()

	for {
		_, err := connector.GetRedisClient(ctx, 0)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("redis completed")
}
