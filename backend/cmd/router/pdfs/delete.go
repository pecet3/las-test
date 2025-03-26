package pdfsRouter

import (
	"net/http"
	"os"

	"github.com/pecet3/logger"
)

func (r router) handleDeletePDF(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	uuid := req.PathValue("uuid")
	pdf, err := r.app.Data.GetPdfByUUID(req.Context(), uuid)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if pdf.UserID != u.ID {
		logger.Warn("user wants to delete pdf not belogns to them")
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	dir := r.app.PDF.GetUserUploadDir(u) + "/" + pdf.Name

	if err := os.Remove(dir); err != nil {
		logger.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := r.app.Data.DeletePdf(req.Context(), pdf.ID); err != nil {
		logger.Error(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
