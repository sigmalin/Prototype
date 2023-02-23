package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis"
)

type redisSession struct {
	sid    string
	client *redis.Client
}

func (rs *redisSession) Set(value interface{}) error {
	return rs.client.Set(rs.sid, value, SESSION_DURATION).Err()
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

	_, err2 := rs.client.Expire(rs.sid, SESSION_DURATION).Result()
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

func (rp *redisProvider) redisConnect() {
	conn := fmt.Sprintf("%s:%d", REDIS_ADDRESS, REDIS_PORT)

	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	rp.redisClient = client
}

// init
func initSessionProvider() {

	provider := &redisProvider{}
	provider.redisConnect()

	RegisterSessionProvider("redis", provider)
}
