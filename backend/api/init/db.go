package initial

import (
	"config"
	"log"
	"sync"
	"time"

	connector "connect/db"
)

func initDB(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start DB")

	for {
		db := connector.GetDB(config.SQL_DATABASE)
		if db != nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Printf("DB completed")
}
