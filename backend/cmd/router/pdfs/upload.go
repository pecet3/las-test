package pdfsRouter

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pecet3/logger"
)

const (
	UPLOAD_BASE_DIR = "./uploads/"
)

func (r router) handleUploadPDFs(w http.ResponseWriter, req *http.Request) {
	u, err := r.app.Auth.GetContextUser(req)
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		logger.Error(err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := req.FormFile("file")
	if err != nil {
		logger.Error(err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if filepath.Ext(handler.Filename) != ".pdf" {
		logger.Error(err)
		http.Error(w, "Only PDF files are allowed", http.StatusBadRequest)
		return
	}

	uploadDir := UPLOAD_BASE_DIR + u.FolderUuid

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		logger.Error(err)
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(uploadDir, handler.Filename)

	destFile, err := os.Create(filePath)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Error creating destination file", http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File uploaded successfully"))
}
