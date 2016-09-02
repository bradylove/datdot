package filemanager_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bradylove/dotter/filemanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Init", func() {
	var (
		manager  filemanager.FileManager
		basePath string
		dotDir   string
	)

	BeforeEach(func() {
		basePath = os.TempDir()
		dotDir = filepath.Join(basePath, ".dot")
		os.RemoveAll(dotDir)
		manager = filemanager.New(basePath)

		err := manager.Init("git@github.com:bradylove/make-believe")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(dotDir)
	})

	It("creates the dotter directory", func() {
		fmt.Println(dotDir)
		file, err := os.Open(dotDir)
		Expect(err).ToNot(HaveOccurred())

		stat, err := file.Stat()
		Expect(err).ToNot(HaveOccurred())

		Expect(stat.IsDir()).To(BeTrue())
	})

	It("initializes a git repo", func() {
		file, err := os.Open(filepath.Join(basePath, ".dot", ".git"))
		Expect(err).ToNot(HaveOccurred())

		stat, err := file.Stat()
		Expect(err).ToNot(HaveOccurred())

		Expect(stat.IsDir()).To(BeTrue())
	})

	It("adds a remote origin to the git repo", func() {
		cmd := exec.Command("git", "remote", "show")
		cmd.Dir = dotDir

		out, err := cmd.Output()
		Expect(err).ToNot(HaveOccurred())

		fmt.Println(string(out))

		Expect(string(out)).To(ContainSubstring("origin"))
	})
})
