package cmd

import (
	"os"

	"github.com/bradylove/datdot/filemanager"
)

func newFileManager() filemanager.FileManager {
	return filemanager.New(os.Getenv("HOME"))
}
