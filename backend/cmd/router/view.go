package router

import (
	"net/http"
	"os"

	"github.com/pecet3/las-test-pdf/utils"
	"github.com/pecet3/logger"
)

func (r router) handleView(w http.ResponseWriter, req *http.Request) {
	ip := utils.GetIP(req)
	logger.Info("Serving react files for ip: ", ip)
	fs := http.FileServer(http.Dir(VIEW_DIR))
	path := req.URL.Path
	_, err := os.Stat(VIEW_DIR + path)

	if os.IsNotExist(err) {
		http.ServeFile(w, req, VIEW_DIR+"/index.html")
		return
	}
	w.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	w.Header().Set("ETag", `"some-unique-hash"`)

	fs.ServeHTTP(w, req)
}
