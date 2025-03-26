package register

import (
	"sync"
)

type Register struct {
	sessions sessions
	sMu      sync.RWMutex
}

func New() *Register {
	r := &Register{
		sessions: make(sessions),
	}
	go r.cleanUpExpiredSessions()
	return r
}
