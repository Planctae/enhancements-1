package enhancement_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement"
)

var _ = Describe("Opening an Enhancement Tracking Receipt", func() {
	Describe("OpenReceipt()", func() {
		Context("when the receipt is embedded in the pull request template", func() {
			It("returns an Enhancement Tracking Receipt", func() {
				receipt, err := enhancement.OpenReceipt(filepath.Join("test_fixtures", "example_receipt.md"))
				Expect(err).ToNot(HaveOccurred(), "expected no error when opening embedded tracking receipt test fixture")

				Expect(receipt.Title()).To(Equal(expectedTitle), "expected title to match title from fixture")
				Expect(receipt.SponsoringSIG()).To(Equal(expectedSponsor), "expected sponsoring SIG to match sponsor from fixture")
				Expect(receipt.KEPLocation()).To(Equal(expectedKEPLocation), "expected KEP location to match location from fixture")
				Expect(receipt.ReleaseNoteLocation()).To(Equal(expectedReleaseNoteLocation), "expected release note location to match location from fixture")
				Expect(receipt.ContactEmail()).To(Equal(expectedContactEmail), "expected contact email to match email from fixture")
				Expect(receipt.CurrentMaturity()).To(Equal(expectedCurrentMaturity), "expected current maturity to match maturity from fixture")
				Expect(receipt.TargetMaturity()).To(Equal(expectedTargetMaturity), "expected target maturity to match maturity from fixture")

			})
		})

		Context("when the receipt is contained in a standalone YAML file", func() {
			It("returns an Enhancement Tracking Receipt", func() {
				receipt, err := enhancement.OpenReceipt(filepath.Join("test_fixtures", "example_receipt.yaml"))
				Expect(err).ToNot(HaveOccurred(), "expected no error when opening standalone tracking receipt test fixture")

				Expect(receipt.Title()).To(Equal(expectedTitle), "expected title to match title from fixture")
				Expect(receipt.SponsoringSIG()).To(Equal(expectedSponsor), "expected sponsoring SIG to match sponsor from fixture")
				Expect(receipt.KEPLocation()).To(Equal(expectedKEPLocation), "expected KEP location to match location from fixture")
				Expect(receipt.ReleaseNoteLocation()).To(Equal(expectedReleaseNoteLocation), "expected release note location to match location from fixture")
				Expect(receipt.ContactEmail()).To(Equal(expectedContactEmail), "expected contact email to match email from fixture")
				Expect(receipt.CurrentMaturity()).To(Equal(expectedCurrentMaturity), "expected current maturity to match maturity from fixture")
				Expect(receipt.TargetMaturity()).To(Equal(expectedTargetMaturity), "expected target maturity to match maturity from fixture")

			})
		})
	})
})
