package filemanager

import (
	"os"
)

func (m *FileManager) Init() error {
	return os.Mkdir(m.dirPath, os.ModeDir|0664)
}
