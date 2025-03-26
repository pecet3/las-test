package auth

import (
	"context"
	"time"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/pkg/auth/login"
	"github.com/pecet3/las-test-pdf/pkg/auth/register"
	"github.com/pecet3/logger"
)

const (
	POST_SUSPEND_PUNISH_DURATION = time.Second * 10
	POST_SUSPEND_DURATION        = time.Second * 1
	AUTH_SESSION_EXPIRY          = time.Hour * 168 * 1 // one week
)

type Mailable interface {
	Send(ctx context.Context, to, code, userName string) error
}

type Sessionable interface {
	Get(string) (interface{}, string, error)
	Add(interface{}) error
	Remove(string)
}

type CheckRole = func(uID int64) bool

type Auth struct {
	d        *data.Queries
	Mailers  map[string]Mailable
	Sessions map[string]Sessionable
	New      makers
	roles    map[string]CheckRole
}

func New(d *data.Queries) *Auth {
	a := &Auth{
		d:        d,
		Mailers:  make(map[string]Mailable),
		Sessions: make(map[string]Sessionable),
		roles:    make(map[string]CheckRole),
	}
	a.Sessions["register"] = register.New()
	a.Sessions["login"] = login.New()

	// a.Mailers["register"] = register.RegisterMailer{}
	// a.Mailers["login"] = login.LoginMailer{}

	a.Mailers["register"] = testMailer{}
	a.Mailers["login"] = testMailer{}

	return a
}

func (a *Auth) AddRole(name string, checkRole CheckRole) {
	a.roles[name] = checkRole
}

type testMailer struct{}

func (testMailer) Send(ctx context.Context, to, code, userName string) error {
	logger.Debug("test email send", to, code, userName)
	return nil
}
