package filemanager_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bradylove/datdot/filemanager"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Init", func() {
	var (
		basePath string
		dotDir   string

		testRepo = "git@github.com:bradylove/datdot-test.git"
	)

	BeforeEach(func() {
		basePath = os.TempDir()
		dotDir = filepath.Join(basePath, ".dot")
		os.RemoveAll(dotDir)

		err := filemanager.New(basePath).Init(testRepo)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(dotDir)
	})

	It("creates the datdot directory", func() {
		file, err := os.Open(dotDir)
		Expect(err).ToNot(HaveOccurred())

		stat, err := file.Stat()
		Expect(err).ToNot(HaveOccurred())

		Expect(stat.IsDir()).To(BeTrue())
	})

	It("writes the initial config file", func() {
		data, err := ioutil.ReadFile(filepath.Join(dotDir, "datdot.json"))
		Expect(err).ToNot(HaveOccurred())

		Expect(string(data)).To(MatchJSON(fmt.Sprintf(`{
			"remote": "%s"
		}`, testRepo)))
	})

	It("initializes a git repo", func() {
		file, err := os.Open(filepath.Join(basePath, ".dot", ".git"))
		Expect(err).ToNot(HaveOccurred())

		stat, err := file.Stat()
		Expect(err).ToNot(HaveOccurred())

		Expect(stat.IsDir()).To(BeTrue())
	})

	It("adds a remote origin to the git repo", func() {
		cmd := exec.Command("git", "remote", "-v")
		cmd.Dir = dotDir

		out, err := cmd.Output()
		Expect(err).ToNot(HaveOccurred())

		Expect(string(out)).To(ContainSubstring("origin"))
		Expect(string(out)).To(ContainSubstring(testRepo))
	})
})
