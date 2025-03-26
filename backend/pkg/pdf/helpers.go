package pdf

import "github.com/pecet3/las-test-pdf/data"

func (PDF) GetUserUploadDir(u data.User) string {
	return UPLOADS_BASE_DIR + "/" + u.FolderUuid
}

func getUserUploadURL(u data.User) string {
	return UPLOADS_BASE_URL + "/" + u.Uuid
}
