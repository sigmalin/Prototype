package initial

import (
	"config"
	"log"
	"sync"
	"time"

	connector "connect/redis"
	"session"
	provider "session/provider/redis"
)

func initSession(wg *sync.WaitGroup) {

	defer wg.Done()

	log.Printf("start session")

	for {
		client, err := connector.GetRedisClient(0)
		if err == nil {
			session.RegisterSessionProvider(config.SESSION_PROVIDER_KEY, provider.NewProvider(client))
			break
		}

		time.Sleep(1 * time.Second)
	}

	mgr, _ := session.NewSessionManager(config.SESSION_PROVIDER_KEY, config.SESSION_NAME)
	session.RegisterSessionManager(config.SESSION_MANAGER_KEY, mgr)

	go mgr.GC(config.SESSION_EXPIRATION)

	log.Printf("session completed")
}
