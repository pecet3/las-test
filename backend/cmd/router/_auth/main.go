package authRouter

import "github.com/pecet3/las-test-pdf/cmd/repos"

type router struct {
	app *repos.App
}

const (
	PREFIX = "/api/auth"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX
)

func Run(
	app *repos.App,
) {

	r := router{
		app: app,
	}
	app.Srv.HandleFunc(POST+"/register", r.handleRegister)
	app.Srv.HandleFunc(POST+"/register/exchange", r.handleRegisterExchange)

	app.Srv.HandleFunc(POST+"/login", r.handleLogin)
	app.Srv.HandleFunc(POST+"/login/exchange", r.handleLoginExchange)

	app.Srv.Handle(POST+"/change-email", r.app.Auth.Authorize(r.handleChangeEmail))
	app.Srv.Handle(POST+"/change-email/exchange", r.app.Auth.Authorize(r.handleChangeEmailExchange))

	app.Srv.Handle(GET+"/logout", r.app.Auth.Authorize(r.handleLogout))
	app.Srv.Handle(GET+"/ping", r.app.Auth.Authorize(r.handlePing))
}
