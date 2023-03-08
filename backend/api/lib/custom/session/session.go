package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session interface {
	Set(value interface{}) error
	Get() interface{}
	Delete() error
	SessionID() string
}

type SessionProvider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(expires int)
}

type SessionManager struct {
	cookieName string
	lock       sync.Mutex
	provider   SessionProvider
}

func (mgr *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		sid := mgr.sessionId()
		session, _ = mgr.provider.SessionInit(sid)
		cookie := http.Cookie{Name: mgr.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: 0}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = mgr.provider.SessionRead(sid)
	}
	return
}

func (mgr *SessionManager) SessionRead(w http.ResponseWriter, r *http.Request) (Session, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		return nil, err
	}

	sid, _ := url.QueryUnescape(cookie.Value)
	session, err := mgr.provider.SessionRead(sid)
	if err != nil {
		expiration := time.Now()
		cookie := http.Cookie{Name: mgr.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
		return nil, err
	}
	return session, nil
}

func (mgr *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		mgr.lock.Lock()
		defer mgr.lock.Unlock()

		mgr.provider.SessionDestroy(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: mgr.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (mgr *SessionManager) GC(expires int) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.provider.SessionGC(expires)
	time.AfterFunc(time.Duration(expires)*time.Second, func() { mgr.GC(expires) })
}

func (mgr *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	/*
		if _, err := rand.Read(b); err != nil {
			return ""
		}
	*/
	return base64.URLEncoding.EncodeToString(b)
}

var sessionProvides = make(map[string]SessionProvider)

func RegisterSessionProvider(key string, provider SessionProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := sessionProvides[key]; dup {
		panic("session: Register called twice for provider " + key)
	}
	sessionProvides[key] = provider
}

func CreateSessionManager(provideName string, cookieName string) (*SessionManager, error) {
	provider, ok := sessionProvides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &SessionManager{cookieName: cookieName, provider: provider}, nil
}
