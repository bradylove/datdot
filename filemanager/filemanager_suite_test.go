package filemanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"os"

	"github.com/bradylove/dotter/filemanager"
)

func TestFilemanager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Filemanager Suite")
}

var _ = BeforeSuite(func() {
	filemanager.Stdout = GinkgoWriter
	filemanager.Stderr = GinkgoWriter
})

var _ = AfterSuite(func() {
	filemanager.Stderr = os.Stdout
	filemanager.Stderr = os.Stderr
})
