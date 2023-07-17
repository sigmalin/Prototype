package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session interface {
	Set(ctx context.Context, value interface{}) error
	Get(ctx context.Context) interface{}
	Delete(ctx context.Context) error
	SessionID() string
}

type SessionProvider interface {
	SessionInit(ctx context.Context, sid string) (Session, error)
	SessionRead(ctx context.Context, sid string) (Session, error)
	SessionDestroy(ctx context.Context, sid string) error
	SessionGC(ctx context.Context, expires int)
}

type SessionManager struct {
	cookieName string
	lock       sync.Mutex
	provider   SessionProvider
}

func (mgr *SessionManager) SessionStart(ctx context.Context, w http.ResponseWriter, r *http.Request) (session Session) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		sid := mgr.sessionId()
		session, _ = mgr.provider.SessionInit(ctx, sid)
		cookie := http.Cookie{Name: mgr.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: 0}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = mgr.provider.SessionRead(ctx, sid)
	}
	return
}

func (mgr *SessionManager) SessionRead(ctx context.Context, w http.ResponseWriter, r *http.Request) (Session, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		return nil, err
	}

	sid, _ := url.QueryUnescape(cookie.Value)
	session, err := mgr.provider.SessionRead(ctx, sid)
	if err != nil {
		expiration := time.Now()
		cookie := http.Cookie{Name: mgr.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
		return nil, err
	}
	return session, nil
}

func (mgr *SessionManager) SessionDestroy(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		mgr.lock.Lock()
		defer mgr.lock.Unlock()

		mgr.provider.SessionDestroy(ctx, cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: mgr.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (mgr *SessionManager) GC(ctx context.Context, expires int) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.provider.SessionGC(ctx, expires)
	time.AfterFunc(time.Duration(expires)*time.Second, func() { mgr.GC(ctx, expires) })
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
