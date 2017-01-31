package filemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bradylove/dotter/config"
	"github.com/spf13/viper"
)

func (m *FileManager) Init(remote string) error {
	if err := cloneRemote(m.dotDirPath, remote); err != nil {
		fmt.Println(err)

		if err := makeDir(m.dotDirPath); err != nil {
			return err
		}

		return initGitRepo(m.dotDirPath, remote)
	}

	return writeConfig(m.dotDirPath, remote)
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

func writeConfig(local, remote string) error {
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	cfg.Remote = remote

	json, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	configPath := filepath.Join(local, "dotter.json")

	return ioutil.WriteFile(configPath, json, os.ModePerm)
}
