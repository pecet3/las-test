package login

import (
	"errors"
	"fmt"
	"time"

	"github.com/pecet3/las-test-pdf/utils"
	"github.com/pecet3/logger"
)

const (
	EXPIRY_DURATION  = time.Second * 60 * 2
	BLOCK_DURATION   = time.Second * 60 * 60
	CLEANUP_DURATION = time.Second * 60 * 60
)

type LoginSession struct {
	ActivateCode    string
	UUID            string
	Expiry          time.Time
	IsBlocked       bool
	UserID          int64
	AttemptCounter  int
	ExchangeCounter int
}

type sessions = map[string]*LoginSession

func NewSession(
	uID int64,
	uuid string) (*LoginSession, string) {
	expiresAt := time.Now().Add(EXPIRY_DURATION)

	code := utils.GenerateCode()
	ea := &LoginSession{
		Expiry:       expiresAt,
		ActivateCode: code,
		UserID:       uID,
		UUID:         uuid,
	}
	return ea, code
}

func (ec *Login) Get(uuid string) (interface{}, string, error) {
	ec.sMu.Lock()
	defer ec.sMu.Unlock()
	session, exists := ec.sessions[uuid]
	if !exists {
		return nil, "", errors.New("session doesn't exist")
	}
	if time.Now().After(session.Expiry) {
		return nil, "", errors.New("Time to enter the code has passed. Try again by resending the email.")
	}
	if session.IsBlocked {
		errmsg := fmt.Sprintf(`Account is blocked, left: %d minutes`, int(session.Expiry.Sub(time.Now()).Minutes()))
		return nil, "", errors.New(errmsg)
	}
	if session.ExchangeCounter >= 5 {
		session.IsBlocked = true
		session.Expiry = time.Now().Add(BLOCK_DURATION)

		go ec.removeBlockedAfterDuration(uuid)
		errmsg := fmt.Sprintf(`Too many attempts, Your account has been blocked for an hour`)
		return nil, "", errors.New(errmsg)
	}

	session.ExchangeCounter += 1
	return session, session.ActivateCode, nil
}
func (ss *Login) Add(s interface{}) error {
	session, _ := s.(*LoginSession)
	ss.sMu.Lock()
	defer ss.sMu.Unlock()
	es, exists := ss.sessions[session.UUID]
	if !exists {
		ss.sessions[session.UUID] = session
		return nil
	}
	if es.IsBlocked {
		errmsg := fmt.Sprintf(`Account is blocked, left:%d minutes`, int(es.Expiry.Sub(time.Now()).Minutes()))
		return errors.New(errmsg)
	}
	if es.AttemptCounter >= 5 {
		es.IsBlocked = true
		es.Expiry = time.Now().Add(BLOCK_DURATION)

		go ss.removeBlockedAfterDuration(session.UUID)
		errmsg := fmt.Sprintf(`Too many attempts, Your account has been blocked for an hour`)
		return errors.New(errmsg)
	}
	es.AttemptCounter = es.AttemptCounter + 1
	es.ActivateCode = session.ActivateCode
	es.Expiry = session.Expiry
	return nil
}

func (ss *Login) Remove(uuid string) {
	ss.sMu.Lock()
	defer ss.sMu.Unlock()
	delete(ss.sessions, uuid)
}

func (ss *Login) cleanUpExpiredSessions() {
	for {
		time.Sleep(CLEANUP_DURATION)
		cleanedSessions := 0
		ss.sMu.Lock()
		for token, session := range ss.sessions {
			if time.Now().After(session.Expiry) {
				delete(ss.sessions, token)
				cleanedSessions++
			}
		}
		ss.sMu.Unlock()
		logger.Info(fmt.Sprintf(`Cleaned Expired Login Sessions: %d`, cleanedSessions))
	}
}

func (ss *Login) removeBlockedAfterDuration(uuid string) {
	time.Sleep(BLOCK_DURATION)
	ss.sMu.Lock()
	delete(ss.sessions, uuid)
	logger.InfoC("Removed blocked user with uuid", uuid)
	ss.sMu.Unlock()
}
