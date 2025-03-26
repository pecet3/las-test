package pdf

import (
	"github.com/pecet3/las-test-pdf/data"
)

func (PDF) GetPdfURL(u data.User, fileName string) string {
	uDir := getUserUploadURL(u)
	return uDir + "/?name=" + fileName
}
