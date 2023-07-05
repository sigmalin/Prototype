package initial

import (
	"log"
	"sync"
)

func AllService() {

	log.Printf("start initial")

	wg := &sync.WaitGroup{}

	wg.Add(3)

	initSession(wg)

	initDB(wg)

	initUID(wg)

	wg.Wait()

	log.Printf("initial completed")
}
