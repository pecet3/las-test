package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/utils"
	"github.com/pecet3/logger"
)

type contextKey string

const sessionContextKey contextKey = "session"

func (as *Auth) getTokens(r *http.Request) (string, string, error) {
	var jwt string
	var refresh string

	cookie, err := r.Cookie("auth")
	if err != nil || cookie.Value == "" {
		return "", "", err
	}
	rCookie, err := r.Cookie("refresh")
	if err != nil || cookie.Value == "" {
		return "", "", err
	}
	refresh = rCookie.Value
	jwt = cookie.Value

	if jwt == "" {
		return "", "", errors.New("no jwt")
	}

	return jwt, refresh, nil
}

func (as *Auth) checkRoles(uID int64, roles ...string) bool {
	for _, r := range roles {
		check, ok := as.roles[r]
		if !ok {
			return false
		}
		if !check(uID) {
			return false
		}
	}
	return true
}

func (as *Auth) Authorize(next http.HandlerFunc, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, refresh, err := as.getTokens(r)
		if err != nil {
			logger.Error(err.Error()+" IP:", utils.GetIP(r))
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		_, err = validateJWT(jwt)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		s, err := as.GetAuthSession(jwt)
		if err != nil || s.Token == "" {
			logger.Warn("<Auth> Session doesn't exist")
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		// check roles

		if !as.checkRoles(s.UserID, roles...) {
			logger.Warn("invalid permissions uID: ", s.UserID)
			http.Error(w, "invalid permissions", http.StatusForbidden)
			return
		}

		logger.Info(fmt.Sprintf(`%s %s uID: %d`, r.Method, r.RequestURI, s.UserID))
		if s.IsExpired.Bool {
			logger.Warn("expired session with uID: ", s.UserID)
			http.Error(w, "expired session", http.StatusUnauthorized)
			return
		}
		if s.ActivateCode != jwt {
			logger.WarnC("invalid jwt token. user id: ", s.UserID)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		if s.Expiry.Before(time.Now()) {
			if refresh != s.RefreshToken {
				http.Error(w, "Your sessions is expired you need to login once again", http.StatusUnauthorized)
				if err := as.UpdateSessionIsExpired(s.Token); err != nil {
					logger.Error(err)
					http.Error(w, "", http.StatusUnauthorized)
					return
				}
				return
			}
		}
		// PostSuspend is for methods post / put / delete

		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete {
			if !s.PostSuspendExpiry.Time.IsZero() && !s.PostSuspendExpiry.Time.Before(time.Now()) {
				logger.Warn(fmt.Sprintf(`<Auth> User with ID: %d is trying to use method %v, but they are  suspended`, s.UserID, r.Method))

				http.Error(w, fmt.Sprintf(`You are suspended for %v seconds to avoid spam`,
					POST_SUSPEND_PUNISH_DURATION.Seconds()), http.StatusBadRequest)
				if err := as.d.UpdateSessionPostSuspendExpiry(
					r.Context(),
					data.UpdateSessionPostSuspendExpiryParams{
						Token:             s.Token,
						PostSuspendExpiry: sql.NullTime{Time: time.Now().Add(POST_SUSPEND_PUNISH_DURATION), Valid: true}}); err != nil {
					logger.Error(err)
					http.Error(w, "", http.StatusUnauthorized)
					return
				}
				return
			}

			err := as.d.UpdateSessionPostSuspendExpiry(
				r.Context(),
				data.UpdateSessionPostSuspendExpiryParams{
					Token:             s.Token,
					PostSuspendExpiry: sql.NullTime{Time: time.Now().Add(POST_SUSPEND_DURATION), Valid: true}})

			if err != nil {
				logger.Error(err)
				http.Error(w, "", http.StatusUnauthorized)
				return
			}

		}
		ctx := context.WithValue(r.Context(), sessionContextKey, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
