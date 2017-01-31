package filemanager

import (
	"fmt"
	"os"
)

func (m *FileManager) Init(remote string) error {
	if err := cloneRemote(m.dotDirPath, remote); err != nil {
		fmt.Println(err)

		if err := makeDir(m.dotDirPath); err != nil {
			return err
		}

		return initGitRepo(m.dotDirPath, remote)
	}

	return m.setConfigRemote(remote)
}

func cloneRemote(local, remote string) error {
	return run(os.Getenv("HOME"), "git", "clone", remote, local)
}

func makeDir(path string) error {
	return os.Mkdir(path, os.ModeDir|0775)
}

func initGitRepo(local, remote string) error {
	if err := run(local, "git", "init"); err != nil {
		return err
	}

	return run(local, "git", "remote", "add", "origin", remote)
}
