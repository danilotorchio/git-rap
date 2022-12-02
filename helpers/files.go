package helpers

import (
	"io"
	"os"
)

func CheckIfDirectoryExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CheckIfDirectoryIsEmpty(path string) bool {
	d, err := os.Open(path)
	CheckIfError(err)
	defer d.Close()

	_, err = d.Readdirnames(1)
	return err == io.EOF
}

func CreateDirectory(path string) {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(path, 0755)
		} else {
			CheckIfError(err)
		}
	}
}
