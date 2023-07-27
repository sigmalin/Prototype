package initial

import (
	"log"
	"sync"

	"config"
	"uid"
	satori "uid/satori"
)

func initUID(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start uid")

	uid.Register(config.UID_GENERATOR_KEY, satori.Generate)

	log.Printf("uid completed")
}

func releaseUID() {

}
