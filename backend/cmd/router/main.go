package router

import "github.com/pecet3/las-test-pdf/cmd/repos"

const (
	PREFIX = "/api"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX

	IMAGES_DIR = "./cmd/images"
	VIEW_DIR   = "./cmd/view"
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

	app.Srv.HandleFunc("/", r.handleView)

}
