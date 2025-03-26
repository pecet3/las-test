package pdf

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pecet3/las-test-pdf/data"
)

const MAX_KB = 500 * 1024

func (p PDF) SavePdfFromRequest(req *http.Request, u data.User) (string, error) {
	if err := req.ParseMultipartForm(MAX_KB); err != nil {
		return "", err
	}
	file, handler, err := req.FormFile("file")

	if handler.Size < MAX_KB {
		return "", errors.New("too large file")
	}
	if err != nil {
		return "", err
	}
	defer file.Close()

	if filepath.Ext(handler.Filename) != ".pdf" {
		return "", err
	}

	uploadDir := p.GetUserUploadDir(u)

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, handler.Filename)

	destFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}
	return handler.Filename, nil
}
