package repos

import (
	"net/http"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/data/dtos"
)

type App struct {
	Srv  *http.ServeMux
	Data *data.Queries
	Dto  dtos.Dto
}

func NewApp() *App {
	mux := http.NewServeMux()
	db := data.NewSQLite()
	data := data.New(db)

	return &App{
		Srv:  mux,
		Data: data,
		Dto:  dtos.New(),
	}

}
