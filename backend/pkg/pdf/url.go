package pdf

import (
	"os"

	"github.com/pecet3/las-test-pdf/data"
)

func (PDF) GetPdfURL(u data.User, fileName string) string {
	uDir := getUserUploadURL(u)
	addr := os.Getenv("ADDRESS")
	return addr + uDir + "/" + fileName
}
