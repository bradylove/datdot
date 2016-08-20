package filemanager

import (
	"os"
	"os/exec"
)

func (m *FileManager) Init() error {
	if err := makeDir(m.dirPath); err != nil {
		return err
	}

	return initGitRepo(m.dirPath)
}

func makeDir(path string) error {
	return os.Mkdir(path, os.ModeDir|0775)
}

func initGitRepo(path string) error {
	cmd := exec.Command("git", "init", path)

	return cmd.Run()
}
