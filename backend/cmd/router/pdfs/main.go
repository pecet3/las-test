package pdfsRouter

import "github.com/pecet3/las-test-pdf/cmd/repos"

type router struct {
	app *repos.App
}

const (
	PREFIX = "/api/pdfs"
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

	app.Srv.Handle(POST, r.app.Auth.Authorize(r.handleUploadPDFs))
	app.Srv.Handle(GET, r.app.Auth.Authorize(r.handleGetUserPDFsDto))
}
