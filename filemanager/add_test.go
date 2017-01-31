package filemanager_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/bradylove/dotter/filemanager"
	"github.com/bradylove/dotter/testhelpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add", func() {
	var (
		manager *filemanager.FileManager

		basePath     string
		testFileName string
		testFilePath string

		testFileModTime time.Time

		testRepo = "git@github.com:bradylove/dotter-test.git"
	)

	BeforeEach(func() {
		testhelpers.Clean()
		testhelpers.Prepare()

		basePath = testhelpers.Helper.FakeHome

		testFileName = "test-file"
		testFilePath = filepath.Join(basePath, testFileName)
		testFileModTime = time.Now().Add(-time.Hour)

		testhelpers.Helper.CreateFile(testFileName, "hello-world\n", 0755)
		Expect(os.Chtimes(testFilePath, testFileModTime, testFileModTime))

		manager = filemanager.New(basePath)
		Expect(manager.Init(testRepo)).To(Succeed())

		testhelpers.InitViper(basePath)
	})

	AfterEach(func() {
		testhelpers.Helper.Clean()
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

		It("adds the dotfile to the config", func() {
			data, err := ioutil.ReadFile(filepath.Join(basePath, ".dot", "dotter.json"))
			Expect(err).ToNot(HaveOccurred())

			Expect(string(data)).To(MatchJSON(fmt.Sprintf(`{
				"remote": "%s",
				"dotfiles": {
					"%s": "$HOME/%s"
				}
			}`, testRepo, testFileName, testFileName)))
		})
	})

	Context("when adding a directory", func() {
		BeforeEach(func() {
			testhelpers.Helper.CreateFile("nested/dir/nested-file.txt", "hello-world\n", 0666)

			Expect(manager.Add(filepath.Join(basePath, "nested"))).To(Succeed())
		})

		It("adds the whole directory", func() {
			dirStat, err := os.Stat(filepath.Join(basePath, ".dot", "nested", "dir"))
			Expect(err).ToNot(HaveOccurred())

			Expect(dirStat.Mode().Perm()).To(Equal(os.FileMode(0755)))
			Expect(dirStat.IsDir()).To(BeTrue())

			_, err = os.Stat(filepath.Join(basePath, ".dot", "nested", "dir", "nested-file.txt"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("adds the directory to the config", func() {
			data, err := ioutil.ReadFile(filepath.Join(basePath, ".dot", "dotter.json"))
			Expect(err).ToNot(HaveOccurred())

			Expect(string(data)).To(MatchJSON(fmt.Sprintf(`{
				"remote": "%s",
				"dotfiles": {
					"nested": "$HOME/nested"
				}
			}`, testRepo)))
		})
	})

	Context("when adding a file nested in directories", func() {
		BeforeEach(func() {
			nestedFile := filepath.Join(basePath, "nested", "dir", "nested-file.txt")
			testhelpers.Helper.CreateFile("nested/dir/nested-file.txt", "hello-world\n", 0666)

			Expect(manager.Add(nestedFile)).To(Succeed())
		})

		It("creates the nested structure", func() {
			_, err := os.Stat(filepath.Join(basePath, ".dot", "nested", "dir", "nested-file.txt"))
			Expect(err).ToNot(HaveOccurred())
		})

		It("adds the dotfile to the config", func() {
			data, err := ioutil.ReadFile(filepath.Join(basePath, ".dot", "dotter.json"))
			Expect(err).ToNot(HaveOccurred())

			Expect(string(data)).To(MatchJSON(fmt.Sprintf(`{
				"remote": "%s",
				"dotfiles": {
					"nested/dir/nested-file.txt": "$HOME/nested/dir/nested-file.txt"
				}
			}`, testRepo)))
		})
	})

	Context("when the destination file already exists", func() {
		BeforeEach(func() {
			testFilePath = filepath.Join(basePath, "existing")
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
