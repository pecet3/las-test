package pdfsRouter

import (
	"net/http"

	"github.com/pecet3/las-test-pdf/data/dtos"
	"github.com/pecet3/logger"
)

func (r router) handleGetUserPDFsDto(w http.ResponseWriter, req *http.Request) {
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

	var dto []dtos.PDF

	for _, pdf := range pdfs {
		url := r.app.PDF.GetPdfURL(u, pdf.Name)
		dto = append(dto, *pdf.ToDto(r.app.Data, url))
	}

	if err := r.app.Dto.Send(w, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
