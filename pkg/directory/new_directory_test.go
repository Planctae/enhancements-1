package directory_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/directory"
)

var _ = Describe("creating a new release tracking directory", func() {
	Describe("New()", func() {
		It("creates a new release tracking directory with sub directories representing tracking outcome", func() {
			containingDir, err := ioutil.TempDir("", "reltrackr-dir-test")
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a temp dir to contain the test release dir")
			defer os.RemoveAll(containingDir)

			releaseName := "release-1.15"

			err = directory.New(containingDir, releaseName)
			Expect(err).ToNot(HaveOccurred(), "expected no error when creating a `release-1.15` directory")

			releaseDir := filepath.Join(containingDir, releaseName)
			Expect(releaseDir).To(BeADirectory(), "expected a directory with the release name to be created")

			Expect(filepath.Join(releaseDir, "proposed")).To(BeADirectory(), "expected a `proposed` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "accepted")).To(BeADirectory(), "expected an `accepted` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "rejected")).To(BeADirectory(), "expected a `rejected` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "at-risk")).To(BeADirectory(), "expected an `at-risk` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "shipped")).To(BeADirectory(), "expected a `shipped` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "slipped")).To(BeADirectory(), "expected a `slipped` directory to be contained within the release directory")
			Expect(filepath.Join(releaseDir, "OWNERS")).To(BeARegularFile(), "expected an `OWNERS` file to be contained wtihin the release release directory")

		})
	})
})
