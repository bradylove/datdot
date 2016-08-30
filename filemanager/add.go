package filemanager

import (
	"errors"
	"io"
	"os"
	"path/filepath"
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
	dst := filepath.Join(m.dirPath, filename)

	if err := safeCopyFile(dst, abs); err != nil {
		return err
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return makeSymlink(path, dst)
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

func makeSymlink(dst, src string) error {
	return os.Symlink(src, dst)
}
