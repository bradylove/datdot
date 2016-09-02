package filemanager

import "os"

func (m *FileManager) Init(remote string) error {
	if err := makeDir(m.dirPath); err != nil {
		return err
	}

	return initGitRepo(m.dirPath, remote)
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
