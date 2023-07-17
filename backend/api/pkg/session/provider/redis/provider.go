package sredis

import (
	"context"
	"session"
	"sync"

	"github.com/go-redis/redis/v8"
)

type redisProvider struct {
	lock        sync.Mutex
	redisClient *redis.Client
}

func (rp *redisProvider) SessionInit(ctx context.Context, sid string) (session.Session, error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session, nil
}

func (rp *redisProvider) SessionRead(ctx context.Context, sid string) (session.Session, error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session, session.Update(ctx)
}

func (rp *redisProvider) SessionDestroy(ctx context.Context, sid string) error {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session.Delete(ctx)
}

func (rp *redisProvider) SessionGC(ctx context.Context, expires int) {

}
