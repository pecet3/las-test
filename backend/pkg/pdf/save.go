package pdf

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pecet3/las-test-pdf/data"
)

const MAX_KB = 500 * 1024

func (p PDF) SavePdfFromRequest(req *http.Request, u data.User) (string, error) {
	if err := req.ParseMultipartForm(MAX_KB); err != nil {
		return "", err
	}
	file, handler, err := req.FormFile("file")

	if handler.Size > MAX_KB {
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

	fileName := handler.Filename
	baseName := strings.TrimSuffix(fileName, ".pdf")
	ext := ".pdf"
	counter := 1
	for {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			break
		}
		newFileName := fmt.Sprintf("%s(%d)%s", baseName, counter, ext)
		filePath = filepath.Join(uploadDir, newFileName)
		fileName = newFileName
		counter++
	}

	destFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
