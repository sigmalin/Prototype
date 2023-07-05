package session

import "fmt"

var sessionProvides = make(map[string]SessionProvider)

var sessionManagers = make(map[string]*SessionManager)

func RegisterSessionProvider(key string, provider SessionProvider) {
	if provider == nil {
		panic("session: cannot register provider with nil value")
	}
	if _, dup := sessionProvides[key]; dup {
		panic(fmt.Errorf("session: cannot register the same provider %s", key))
	}
	sessionProvides[key] = provider
}

func RegisterSessionManager(key string, mgr *SessionManager) {
	if mgr == nil {
		panic("session: cannot register manager with nil value")
	}

	if _, exist := sessionManagers[key]; exist {
		panic(fmt.Errorf("session: cannot register the same manager %s", key))
	}

	sessionManagers[key] = mgr
}

func GetManager(key string) *SessionManager {
	mgr, exist := sessionManagers[key]
	if !exist {
		panic(fmt.Errorf("session: cannot find session manager %s", key))
	}
	return mgr
}
