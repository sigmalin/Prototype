package sredis

import (
	"session"
	"sync"

	"github.com/go-redis/redis"
)

type redisProvider struct {
	lock        sync.Mutex
	redisClient *redis.Client
}

func (rp *redisProvider) SessionInit(sid string) (session.Session, error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session, nil
}

func (rp *redisProvider) SessionRead(sid string) (session.Session, error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session, session.Update()
}

func (rp *redisProvider) SessionDestroy(sid string) error {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session.Delete()
}

func (rp *redisProvider) SessionGC(expires int) {

}
