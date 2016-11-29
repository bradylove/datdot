package filemanager_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/bradylove/dotter/filemanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	var (
		manager  filemanager.FileManager
		basePath string

		testFileDir  string
		testFileName string
		testFilePath string

		testFileModTime time.Time
	)

	BeforeEach(func() {
		fakeHome := filepath.Join(os.TempDir(), "fake-home")

		Expect(os.RemoveAll(fakeHome)).To(Succeed())
		Expect(os.Mkdir(fakeHome, os.ModePerm)).To(Succeed())

		basePath = fakeHome
		testFileDir = fakeHome

		testFileName = "test-file"
		testFilePath = filepath.Join(testFileDir, testFileName)
		testFileModTime = time.Now().Add(-time.Hour)

		Expect(ioutil.WriteFile(testFilePath, []byte("hello-world\n"), 0755)).To(Succeed())
		Expect(os.Chtimes(testFilePath, testFileModTime, testFileModTime))

		manager = filemanager.New(basePath)
		Expect(manager.Init("git@github.com:bradylove/make-believe.git")).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(filepath.Join(basePath, ".dot"))).To(Succeed())
		Expect(os.RemoveAll(testFileDir)).To(Succeed())
	})

	Context("when the file does not already exist", func() {
		BeforeEach(func() {
			Expect(manager.Add(testFilePath)).To(Succeed())
		})

		It("moves the file to the dotter directory", func() {
			_, err := os.Stat(filepath.Join(basePath, ".dot", testFileName))
			Expect(err).ToNot(HaveOccurred())
		})

		It("replaces the original file with a symlink", func() {
			_, err := os.Lstat(testFilePath)
			Expect(err).ToNot(HaveOccurred())
		})

		It("maintains file permissions", func() {
			stat, err := os.Stat(filepath.Join(basePath, ".dot", testFileName))
			Expect(err).ToNot(HaveOccurred())

			Expect(stat.Mode()).To(Equal(os.FileMode(0755)))
		})

		It("makes a git commit", func() {
			cmd := exec.Command("git", "log")
			cmd.Dir = filepath.Join(basePath, ".dot")

			output, err := cmd.Output()
			Expect(err).ToNot(HaveOccurred())

			Expect(string(output)).To(ContainSubstring(fmt.Sprintf("Add %s (via dotter)", testFileName)))
		})
	})

	Context("when adding a directory", func() {
		BeforeEach(func() {
			nestedDir := filepath.Join(testFileDir, "nested", "dir")
			Expect(os.MkdirAll(nestedDir, 0755)).To(Succeed())

			nestedFile := filepath.Join(nestedDir, "nested-file.txt")
			Expect(ioutil.WriteFile(nestedFile, []byte("hello-world\n"), 0666)).To(Succeed())

			Expect(manager.Add(filepath.Join(testFileDir, "nested"))).To(Succeed())
		})

		It("adds the whole directory", func() {
			dirStat, err := os.Stat(filepath.Join(basePath, ".dot", "nested", "dir"))
			Expect(err).ToNot(HaveOccurred())

			Expect(dirStat.Mode().Perm()).To(Equal(os.FileMode(0755)))
			Expect(dirStat.IsDir()).To(BeTrue())

			_, err = os.Stat(filepath.Join(basePath, ".dot", "nested", "dir", "nested-file.txt"))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when adding a file nested in directories", func() {
		BeforeEach(func() {
			nestedDir := filepath.Join(testFileDir, "nested", "dir")
			Expect(os.MkdirAll(nestedDir, 0755)).To(Succeed())

			nestedFile := filepath.Join(nestedDir, "nested-file.txt")
			Expect(ioutil.WriteFile(nestedFile, []byte("hello-world\n"), 0756)).To(Succeed())

			Expect(manager.Add(nestedFile)).To(Succeed())
		})

		It("creates the nested structure", func() {
			_, err := os.Stat(filepath.Join(basePath, ".dot", "nested", "dir", "nested-file.txt"))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when the destination file already exists", func() {
		BeforeEach(func() {
			testFilePath = filepath.Join(testFileDir, "existing")
			destinationPath := filepath.Join(basePath, ".dot", "existing")

			Expect(ioutil.WriteFile(testFilePath, []byte("hello-world\n"), 0755)).To(Succeed())
			Expect(ioutil.WriteFile(destinationPath, []byte("hello-world\n"), 0755)).To(Succeed())
		})

		It("does not overwrite existing files", func() {
			Expect(manager.Add(testFilePath)).To(MatchError(filemanager.ErrDestinationExists))
		})

		XIt("overwrites an existing file if force is true")
		XIt("mentions the overwrite in the git commit message")
	})
})
