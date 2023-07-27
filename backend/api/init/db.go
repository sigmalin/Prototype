package initial

import (
	"context"
	"log"
	"sync"
	"time"

	connector "connect/mongo"
)

func initDB(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start DB")

	for {
		err := connector.NewClient(context.TODO())
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("DB completed")
}

func releaseDB() {
	connector.Release(context.TODO())
}
