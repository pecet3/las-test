package pdf_test

import (
	"log"
	"testing"

	"github.com/pecet3/las-test-pdf/data"
	"github.com/pecet3/las-test-pdf/pkg/pdf"
)

func TestGetPdfURL(t *testing.T) {

	u := data.User{
		ID:         1,
		Uuid:       "123e4567-e89b-12d3-a456-426614174000",
		Name:       "John Doe",
		FolderUuid: "23e4567-e89b-12d3-a456-426614174000",
	}

	fileName := "test.pdf"

	result := pdf.New().GetPdfURL(u, fileName)

	log.Println(result)
}
