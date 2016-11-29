package filemanager

import (
	"io"
	"os"
	"path/filepath"
)

var (
	Stdin  io.Writer = os.Stdin
	Stdout io.Writer = os.Stdout
	Stderr io.Writer = os.Stderr
)

type FileManager struct {
	dotDirPath  string
	homeDirPath string
}

func New(base string) FileManager {
	return FileManager{
		dotDirPath:  filepath.Join(base, ".dot"),
		homeDirPath: base,
	}
}
