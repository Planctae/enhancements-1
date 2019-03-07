package settings_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/settings"
)

var _ = Describe("Creating runtime settings", func() {
	Describe("NewRuntime()", func() {
		It("builds a new runtime", func() {
			tmpDir, err := ioutil.TempDir("", "reltrackr-settings-test")
			Expect(err).ToNot(HaveOccurred(), "expected no error creating a temp dir for test")
			defer os.RemoveAll(tmpDir)

			runtime, err := settings.NewRuntime(tmpDir)
			Expect(err).ToNot(HaveOccurred(), "expected no error building new runtime settings with a valid directory")

			Expect(runtime.ReceiptsLocation()).To(Equal(tmpDir), "expected receipts location returned by runtime settings to match that given at creation")
		})

		Context("when the location does not exist", func() {
			It("returns an error", func() {
				_, err := settings.NewRuntime("does not exist")
				Expect(err.Error()).To(ContainSubstring("refusing to create runtime settings with non existent location"), "expected error building new runtime settings with a non existent directory")
			})
		})
	})
})

var _ = Describe("Persisting runtime settings", func() {
	Describe("#Persist()", func() {
		It("persists the runtime to the default location", func() {
			tmpDir, err := ioutil.TempDir("", "reltrackr-settings-test")
			Expect(err).ToNot(HaveOccurred(), "expected no error creating a temp dir for test")
			defer os.RemoveAll(tmpDir)

			runtime, err := settings.NewRuntime(tmpDir)
			Expect(err).ToNot(HaveOccurred(), "expected no error building new runtime settings with a valid directory")

			filename := fmt.Sprintf("%s.yaml", settings.Filename)

			err = runtime.Persist(filepath.Join(tmpDir, filename))
			Expect(err).ToNot(HaveOccurred(), "expected no error when persisting valid runtime settings")

			Expect(filepath.Join(tmpDir, filename)).To(BeARegularFile(), "expected a settings file to be written after persist")

			var got struct {
				ReceiptsLocation string `yaml:"receipts_location"`
			}

			settingsBytes, err := ioutil.ReadFile(filepath.Join(tmpDir, filename))
			Expect(err).ToNot(HaveOccurred(), "expected no error reading the settings file location")

			err = yaml.Unmarshal(settingsBytes, &got)
			Expect(err).ToNot(HaveOccurred(), "expected no error unmarshalling the settings bytes as YAML")

			Expect(got.ReceiptsLocation).To(Equal(tmpDir), "expected read settings value to match runtime settings creation argument")
		})
	})
})
