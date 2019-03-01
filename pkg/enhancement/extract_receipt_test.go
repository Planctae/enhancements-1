package enhancement_test

import (
	"io/ioutil"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement"
)

var _ = Describe("Extracting an Enhancement Receipt from the Markdown Template", func() {
	Describe("ExtractFromMarkdown()", func() {
		It("extracts the enhancement receipt from the Pull Request template presented to contributors", func() {
			populatedTemplateBytes, err := ioutil.ReadFile(filepath.Join("test_fixtures", "example_receipt.md"))
			Expect(err).ToNot(HaveOccurred(), "expected no error to occur when reading the example pull request template test fixture")

			receipt, err := enhancement.ExtractFromMarkdown(populatedTemplateBytes)
			Expect(err).ToNot(HaveOccurred(), "expected no error when extracting the tracking receipt from the pull request template test fixture")

			Expect(receipt.Title()).To(Equal(expectedTitle), "expected title to match title from fixture")
			Expect(receipt.SponsoringSIG()).To(Equal(expectedSponsor), "expected sponsoring SIG to match sponsor from fixture")
			Expect(receipt.KEPLocation()).To(Equal(expectedKEPLocation), "expected KEP location to match location from fixture")
			Expect(receipt.ReleaseNoteLocation()).To(Equal(expectedReleaseNoteLocation), "expected release note location to match location from fixture")
			Expect(receipt.ContactEmail()).To(Equal(expectedContactEmail), "expected contact email to match email from fixture")
			Expect(receipt.CurrentMaturity()).To(Equal(expectedCurrentMaturity), "expected current maturity to match maturity from fixture")
			Expect(receipt.TargetMaturity()).To(Equal(expectedTargetMaturity), "expected target maturity to match maturity from fixture")
		})

		Context("when the enhancement receipt is not contained within a codeblock", func() {
			It("returns an error", func() {
				populatedTemplateBytes, err := ioutil.ReadFile(filepath.Join("test_fixtures", "example_receipt_outside_codeblock.md"))
				Expect(err).ToNot(HaveOccurred(), "expected no error to occur when reading the example pull request template test fixture")

				_, err = enhancement.ExtractFromMarkdown(populatedTemplateBytes)
				Expect(err.Error()).To(ContainSubstring("could not determine receipt location"), "expected error to inform that codeblock containing receipt could not be located")
			})
		})
	})
})
