package utils

import (
	"os"
)

const IMAGES_DIR = "./cmd/images"

func CreateUserFolder(uuid string) error {
	folderPath := "./uploads/" + uuid
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		return err
	}
	return nil
}
