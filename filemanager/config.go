package filemanager

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bradylove/dotter/config"
	"github.com/spf13/viper"
)

func (m *FileManager) loadConfig() (*config.Config, error) {
	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (m *FileManager) setConfigRemote(remote string) error {
	cfg, err := m.loadConfig()
	if err != nil {
		return err
	}

	cfg.Remote = remote

	return m.writeConfig(cfg)
}

func (m *FileManager) addConfigDotfile(src, dst string) error {
	cfg, err := m.loadConfig()
	if err != nil {
		return err
	}

	if cfg.Dotfiles == nil {
		cfg.Dotfiles = make(map[string]string)
	}

	cfg.Dotfiles[src] = strings.Replace(dst, m.homeDirPath, "$HOME", 1)

	return m.writeConfig(cfg)
}

func (m *FileManager) writeConfig(cfg *config.Config) error {
	json, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	configPath := filepath.Join(m.dotDirPath, "dotter.json")

	return ioutil.WriteFile(configPath, json, os.ModePerm)
}
