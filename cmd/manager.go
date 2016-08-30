package cmd

import (
	"os"

	"github.com/bradylove/dotter/filemanager"
)

func newFileManager() filemanager.FileManager {
	return filemanager.New(os.Getenv("HOME"))
}
