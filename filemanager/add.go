package filemanager

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrDestinationExists = errors.New("destination file already exists")
)

func (m *FileManager) Add(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	filename := filepath.Base(path)
	pathFromHome := strings.TrimPrefix(filepath.Dir(abs), m.homeDirPath)
	dst := filepath.Join(m.dotDirPath, pathFromHome, filename)

	if err := createDirIfNeeded(dst, abs); err != nil {
		return err
	}

	if err := copyPath(dst, abs); err != nil {
		return err
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	if err := makeSymlink(path, dst); err != nil {
		return err
	}

	if err := m.addConfigDotfile(filename, path); err != nil {
		return err
	}

	return commitFile(m.dotDirPath, dst, filename)
}

func createDirIfNeeded(dst, src string) error {
	srcDir := filepath.Dir(src)
	dstDir := filepath.Dir(dst)

	srcDirStat, err := os.Stat(srcDir)
	if err != nil {
		return err
	}

	os.MkdirAll(dstDir, srcDirStat.Mode().Perm())

	return nil
}

func copyPath(dst, src string) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return safeCopyDir(dst, src)
	}

	return safeCopyFile(dst, src)
}

func safeCopyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	inStat, err := in.Stat()
	if err != nil {
		return err
	}

	if _, err := os.Stat(dst); err == nil {
		return ErrDestinationExists
	}

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, inStat.Mode().Perm())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)

	// Capture close error
	closeErr := out.Close()

	// Handle copy error
	if err != nil {
		return err
	}

	return closeErr
}

func safeCopyDir(dst, src string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, srcInfo.Mode().Perm()); err != nil {
		return err
	}

	dir, err := os.Open(src)
	if err != nil {
		return err
	}

	objects, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePtr := filepath.Join(src, obj.Name())
		dstFilePtr := filepath.Join(dst, obj.Name())

		if obj.IsDir() {
			if err := safeCopyDir(dstFilePtr, srcFilePtr); err != nil {
				return err
			}

			continue
		}

		if err := safeCopyFile(dstFilePtr, srcFilePtr); err != nil {
			return err
		}
	}

	return nil
}

func makeSymlink(dst, src string) error {
	return os.Symlink(src, dst)
}

func commitFile(repo, dst, filename string) error {
	if err := run(repo, "git", "add", dst); err != nil {
		return err
	}

	return run(repo, "git", "commit", "-m", fmt.Sprintf("Add %s (via datdot)", filename))
}
