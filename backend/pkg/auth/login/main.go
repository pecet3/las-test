package login

import (
	"sync"
)

type Login struct {
	sessions sessions
	sMu      sync.RWMutex
}

func New() *Login {
	l := &Login{
		sessions: make(sessions),
	}
	go l.cleanUpExpiredSessions()
	return l
}
