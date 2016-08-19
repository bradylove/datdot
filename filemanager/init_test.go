package filemanager_test

import (
	"github.com/bradylove/dotter/filemanager"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Init", func() {
	var (
		manager  filemanager.FileManager
		basePath string
	)

	BeforeEach(func() {
		basePath = os.TempDir()
		os.RemoveAll(filepath.Join(basePath, ".dot"))
		manager = filemanager.New(basePath)
	})

	It("creates the dotter directory", func() {
		err := manager.Init()
		Expect(err).ToNot(HaveOccurred())

		file, err := os.Open(filepath.Join(basePath, ".dot"))
		Expect(err).ToNot(HaveOccurred())

		stat, err := file.Stat()
		Expect(err).ToNot(HaveOccurred())

		Expect(stat.IsDir()).To(BeTrue())
	})
})
