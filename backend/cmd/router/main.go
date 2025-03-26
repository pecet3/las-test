package router

import (
	"github.com/pecet3/las-test-pdf/cmd/repos"
	authRouter "github.com/pecet3/las-test-pdf/cmd/router/auth"
	pdfsRouter "github.com/pecet3/las-test-pdf/cmd/router/pdfs"
)

const (
	PREFIX = "/api"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX

	VIEW_DIR = "./cmd/view"
)

type router struct {
	app *repos.App
}

func Run(
	app *repos.App,
) {

	r := router{
		app: app,
	}
	authRouter.Run(app)
	pdfsRouter.Run(app)
	app.Srv.HandleFunc("/", r.handleView)
	app.Srv.HandleFunc("/uploads/{uuid}/", r.servePDF)
}
