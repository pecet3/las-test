package router

import (
	"net/http"

	"github.com/pecet3/logger"
)

func (r router) servePDF(w http.ResponseWriter, req *http.Request) {
	if match := req.Header.Get("If-None-Match"); match == `"some-unique-hash"` {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	uuid := req.PathValue("uuid")
	fName := req.URL.Query().Get("name")
	if uuid == "" || fName == "" {
		http.Error(w, "Missing uuid or name", http.StatusBadRequest)
		return
	}
	user, err := r.app.Data.GetUserByUUID(req.Context(), uuid)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dir := r.app.PDF.GetUserUploadDir(user) + "/" + fName
	http.ServeFile(w, req, dir)
}
