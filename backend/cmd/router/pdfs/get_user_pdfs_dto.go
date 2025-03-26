package pdfsRouter

import (
	"net/http"

	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleGetUserPdfsDto(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	pdfs, err := r.app.Data.GetPdfsByUserID(req.Context(), u.ID)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	dto := []dtos.PDF{}

	for _, pdf := range pdfs {
		dto = append(dto, *pdf.ToDto(r.app.Data))
	}

	w.WriteHeader(http.StatusCreated)
}
