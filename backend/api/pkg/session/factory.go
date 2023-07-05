package session

import "fmt"

func NewSessionManager(provideName string, cookieName string) (*SessionManager, error) {
	provider, ok := sessionProvides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &SessionManager{cookieName: cookieName, provider: provider}, nil
}
