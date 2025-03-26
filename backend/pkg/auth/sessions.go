package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/pecet3/las-test-pdf/data"
)

func (as *Auth) GetAuthSession(token string) (data.Session, error) {
	return as.d.GetSessionByToken(context.Background(), token)
}

func (as *Auth) AddAuthSession(token string, session data.AddSessionParams) (data.Session, error) {
	return as.d.AddSession(context.Background(), session)
}
func (as *Auth) UpdateSessionIsExpired(token string) error {
	return as.d.UpdateSessionIsExpired(context.Background(), data.UpdateSessionIsExpiredParams{
		IsExpired: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
		Token: token,
	})
}

func (as *Auth) GetAuthSessionFromRequest(r *http.Request) (data.Session, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return data.Session{}, err
	}
	sessionToken := cookie.Value

	s, err := as.GetAuthSession(sessionToken)
	if err != nil {
		log.Println("<Auth> Session doesn't exist")
		return data.Session{}, err
	}
	return s, nil
}

func (as *Auth) GetContextSession(r *http.Request) (data.Session, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(data.Session)
	if !ok {
		return data.Session{}, errors.New("session not found in context")
	}
	return session, nil
}

func (as *Auth) GetContextUser(r *http.Request) (data.User, error) {
	ctx := r.Context()
	session, ok := ctx.Value(sessionContextKey).(data.Session)
	if !ok {
		return data.User{}, errors.New("session not found in context")
	}
	u, err := as.d.GetUserByID(ctx, int64(session.UserID))
	if err != nil {
		return data.User{}, err
	}
	return u, nil
}
