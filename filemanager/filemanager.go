package filemanager

import (
	"path/filepath"
)

type FileManager struct {
	dirPath string
}

func New(base string) FileManager {
	return FileManager{dirPath: filepath.Join(base, ".dot")}
}
