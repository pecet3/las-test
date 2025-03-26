package router

import (
	"net/http"
)

func (r router) handleImages(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	w.Header().Set("ETag", `"some-unique-hash"`)

	if match := req.Header.Get("If-None-Match"); match == `"some-unique-hash"` {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	fName := req.PathValue("fname")
	http.ServeFile(w, req, IMAGES_DIR+"/"+fName)
}
