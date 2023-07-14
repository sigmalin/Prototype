package initial

import (
	"log"
	"sync"
	"time"

	connector "connect/redis"
)

func initRedis(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start redis")

	for {
		_, err := connector.GetRedisClient(0)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("redis completed")
}
