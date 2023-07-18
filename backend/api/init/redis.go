package initial

import (
	"log"
	"sync"

	cache "cache/redis"
)

func initRedisCache(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start redis cache")

	cache.GetInstance()

	log.Printf("redis cache completed")
}
