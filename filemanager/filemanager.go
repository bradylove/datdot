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
	dirPath string
}

func New(base string) FileManager {
	return FileManager{dirPath: filepath.Join(base, ".dot")}
}
