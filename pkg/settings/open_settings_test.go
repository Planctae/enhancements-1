package settings_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("reading settings from an existing file", func() {
	Describe("Open()", func() {
		It("returns runtime settings", func() {
			tmpDir, err := ioutil.TempDir("", "reltrackr-open-settings-test")
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp dir for test")
			defer os.RemoveAll(tmpDir)

			tmpFile, err := ioutil.TempFile(tmpDir, "reltrackr-open-settings-test")
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp file for test")
			defer os.Remove(tmpFile.Name())

			var testConfiguration struct {
				ReceiptsLocation
			}

		})

		Context("when the expected configuration values are not found", func() {
			It("returns an error", func() {

			})
		})

		Context("when the expected configuration values are invalid", func() {
			It("returns an error", func() {

			})
		})
	})
})

var _ = Describe("reading settings from the environment", func() {
	Describe("Open()", func() {
		It("returns runtime settings", func() {

		})

		Context("when the configuration values are not present in the environemtn", func() {
			It("returns an error", func() {

			})
		})

	})
})
