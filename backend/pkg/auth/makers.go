package auth

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/pkg/auth/login"
	"github.com/pecet3/las-test-pdf/pkg/auth/register"
)

type makers struct {
}

func (makers) LoginSession(
	uID int64,
	uuid string) (*login.LoginSession, string) {
	return login.NewSession(uID, uuid)
}

func (makers) RegisterSession(
	uID int64,
	email string) (*register.RegisterSession, string) {
	return register.NewSession(uID, email)
}

func (m makers) AuthSession(uId int64, uEmail, uName string) (data.AddSessionParams, string, error) {
	expiresAt := time.Now().Add(AUTH_SESSION_EXPIRY)
	jwtToken, err := generateJWT(uEmail, uName)
	if err != nil {
		return data.AddSessionParams{}, "", err
	}
	ea := data.AddSessionParams{
		Type:              "",
		IsExpired:         sql.NullBool{Bool: false},
		UserIp:            "",
		Token:             jwtToken,
		Expiry:            expiresAt,
		UserID:            uId,
		Email:             uEmail,
		ActivateCode:      jwtToken,
		PostSuspendExpiry: sql.NullTime{Time: time.Now().Add(POST_SUSPEND_DURATION)},
		RefreshToken:      uuid.NewString(),
	}
	return ea, jwtToken, nil
}
