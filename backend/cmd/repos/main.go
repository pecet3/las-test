package repos

import (
	"net/http"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/las-test-pdf/pkg/auth"
	"github.com/pecet3/las-test-pdf/pkg/pdf"
)

type App struct {
	Srv  *http.ServeMux
	Data *data.Queries
	Dto  dtos.Dto
	Auth *auth.Auth
	PDF  *pdf.PDF
}

func NewApp() *App {
	mux := http.NewServeMux()
	db := data.NewSQLite()
	data := data.New(db)
	auth := auth.New(data)
	pdf := pdf.New()
	return &App{
		Srv:  mux,
		Data: data,
		Dto:  dtos.New(),
		Auth: auth,
		PDF:  pdf,
	}

}
