package testhelpers

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

type TestHelper struct {
	FakeHome string
}

func (th *TestHelper) Clean() {
	Expect(os.RemoveAll(th.FakeHome)).To(Succeed())
}

func (th *TestHelper) Prepare() {
	Expect(os.Mkdir(th.FakeHome, os.ModePerm)).To(Succeed())
}

func (th *TestHelper) CreateFile(path, text string, perm int) {
	fullPath := filepath.Join(th.FakeHome, path)
	dir := filepath.Dir(fullPath)

	Expect(os.MkdirAll(dir, os.ModePerm)).To(Succeed())
	Expect(ioutil.WriteFile(fullPath, []byte(text), os.FileMode(perm))).To(Succeed())
}

var Helper *TestHelper

func init() {
	Helper = &TestHelper{
		FakeHome: filepath.Join(os.TempDir(), "fake-home"),
	}
}

func Clean() {
	Helper.Clean()
}

func Prepare() {
	Helper.Prepare()
}

func InitViper(path string) {
	viper.SetConfigName("datdot")
	viper.AddConfigPath(filepath.Join(path, ".dot"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
