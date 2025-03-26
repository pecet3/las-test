package pdfsRouter

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/logger"
)

func (r router) handleUploadPDFs(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	fName, err := r.app.PDF.SavePdfFromRequest(req, u)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	r.app.Data.AddPdf(req.Context(), data.AddPdfParams{
		Uuid:   uuid.NewString(),
		UserID: u.ID,
		Name:   fName,
	})
	url := r.app.PDF.GetPdfURL(u, fName)
	logger.Debug(url)
	w.WriteHeader(http.StatusCreated)
}
