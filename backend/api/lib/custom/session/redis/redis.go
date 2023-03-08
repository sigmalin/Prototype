package redis

import (
	"fmt"
	"log"
	"sync"

	"custom/config"
	connector "custom/connect/redis"
	"custom/session"

	"github.com/go-redis/redis"
)

type Session = session.Session

type redisSession struct {
	sid    string
	client *redis.Client
}

func (rs *redisSession) Set(value interface{}) error {
	return rs.client.Set(rs.sid, value, config.SESSION_DURATION).Err()
}

func (rs *redisSession) Get() interface{} {
	val, err := rs.client.Get(rs.sid).Result()
	if err != nil {
		log.Print(err)
	}
	return val
}

func (rs *redisSession) Delete() error {
	key := rs.sid
	rs.sid = ""
	return rs.client.Del(key).Err()
}

func (rs *redisSession) SessionID() string {
	return rs.sid
}

func (rs *redisSession) Update() error {
	count, err1 := rs.client.Exists(rs.sid).Result()
	if err1 != nil {
		return err1
	}

	if count == 0 {
		return fmt.Errorf("session not exist")
	}

	_, err2 := rs.client.Expire(rs.sid, config.SESSION_DURATION).Result()
	return err2
}

type redisProvider struct {
	lock        sync.Mutex
	redisClient *redis.Client
}

func (rp *redisProvider) SessionInit(sid string) (Session, error) {
	rp.lock.Lock()
	defer rp.lock.Unlock()

	session := &redisSession{sid: sid, client: rp.redisClient}
	return session, nil
}

func (rp *redisProvider) SessionRead(sid string) (Session, error) {
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

// init
func init() {

	client, err := connector.GetRedisClient(0)
	if err != nil {
		log.Printf("connect session redis failure!")
		return
	}

	provider := &redisProvider{redisClient: client}

	session.RegisterSessionProvider("redis", provider)
}
