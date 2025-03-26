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
	logger.Debug(uuid, fName)
	if uuid == "" || fName == "" {
		http.Error(w, "Missing uuid or name", http.StatusBadRequest)
		return
	}
	user, _ := r.app.Data.GetUserByUUID(req.Context(), uuid)

	dir := r.app.PDF.GetUserUploadDir(user) + "/" + fName
	http.ServeFile(w, req, dir)
}
