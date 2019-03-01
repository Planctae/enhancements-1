package enhancement_test

import (
	fuzz "github.com/google/gofuzz"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/maturity"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/enhancementfakes"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement"
)

var _ = Describe("Checking receipt validity", func() {
	Describe("Check()", func() {
		It("ensures the title is given", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.TitleReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no title given"), "expected error message to note missing title when no title given in receipt")

		})

		It("ensures there is at least one author listed", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.AuthorsReturns([]string{})

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no authors given"), "expected error message to note missing authors when not given in receipt")

		})

		It("ensures the sponsoring SIG is set", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.SponsoringSIGReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no sponsoring SIG given"), "expected error message to note missing sponsoring SIG when not given in receipt")

		})

		It("ensures there is at least one affected subproject", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.AffectedSubprojectsReturns([]string{})

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no affected subprojects given"), "expected error message to note missing affected subprojects when not given in receipt")

		})

		It("ensures the KEP URL is provided", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.KEPLocationReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no KEP location given"), "expected error message to note missing KEP URL when not given in receipt")

		})

		It("ensures at least one updated test location is provided", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.TestLocationsReturns([]string{})

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no test locations given"), "expected error message to note missing test locations when not given in receipt")

		})

		It("ensures that a location for drafting the release note exists", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.ReleaseNoteLocationReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no release note location given"), "expected error message to note missing release note location when not given in receipt")

		})

		It("ensures at least one updated documentation location is provided", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.DocumentationLocationsReturns([]string{})

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no documentation locations given"), "expected error message to note missing documentation locations when not given in receipt")

		})

		It("ensures a contact email is listed", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.ContactEmailReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no contact email given"), "expected error message to note missing contact email when not given in receipt")

		})

		It("ensures that the previous maturity level is given", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.CurrentMaturityReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no current maturity level given"), "expected error message to note missing current maturity level when not given in receipt")

			trackingReceipt.CurrentMaturityReturns("invalid")
			err = enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("given state: invalid is not valid"), "expected error message to note invalid maturity level given in receipt")
		})

		It("ensures that the target maturity level is given", func() {
			trackingReceipt := generateValidReceipt()
			trackingReceipt.TargetMaturityReturns("")

			err := enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("no target maturity level given"), "expected error message to note missing target maturity level when not given in receipt")

			trackingReceipt.CurrentMaturityReturns("invalid")
			err = enhancement.CheckReceipt(trackingReceipt)
			Expect(err.Error()).To(ContainSubstring("given state: invalid is not valid"), "expected error message to note invalid target maturiy level given in receipt")
		})

		Context("with a fully filled out receipt", func() {
			It("returns no error", func() {
				trackingReceipt := generateValidReceipt()

				err := enhancement.CheckReceipt(trackingReceipt)
				Expect(err).ToNot(HaveOccurred(), "expected no error when checking a fully filled out enhancement tracking receipt")
			})
		})
	})
})

func generateValidReceipt() *enhancementfakes.FakeTrackingReceipt {
	fakeReceipt := &enhancementfakes.FakeTrackingReceipt{}

	var title string
	var authors []string
	var sponsoringSIG string
	var affectedSubprojects []string
	var kepLocation string
	var testLocations []string
	var releaseNoteLocation string
	var documentationLocations []string
	var contactEmail string

	var currentMaturity string = maturity.ReleaseCandidate // must be one of RC|alpha|beta|GA
	var targetMaturity string = maturity.Alpha             // must be one of RC|alpha|beta|GA

	fuzzer := fuzz.New().NilChance(0)
	fuzzer.Fuzz(&title)
	fuzzer.Fuzz(&authors)
	fuzzer.Fuzz(&sponsoringSIG)
	fuzzer.Fuzz(&affectedSubprojects)
	fuzzer.Fuzz(&kepLocation)
	fuzzer.Fuzz(&testLocations)
	fuzzer.Fuzz(&releaseNoteLocation)
	fuzzer.Fuzz(&documentationLocations)
	fuzzer.Fuzz(&contactEmail)

	fakeReceipt.TitleReturns(title + "ok") // ensure there are at least two characters
	fakeReceipt.AuthorsReturns(authors)
	fakeReceipt.SponsoringSIGReturns(sponsoringSIG + "ok") // ensure there are at least two characters
	fakeReceipt.AffectedSubprojectsReturns(affectedSubprojects)
	fakeReceipt.KEPLocationReturns(kepLocation + "ok") // ensure there are at least two characters
	fakeReceipt.TestLocationsReturns(testLocations)
	fakeReceipt.ReleaseNoteLocationReturns(releaseNoteLocation + "ok") // ensure there are at least two characters
	fakeReceipt.DocumentationLocationsReturns(documentationLocations)
	fakeReceipt.ContactEmailReturns(contactEmail + "ok") // ensure there are at least two characters
	fakeReceipt.CurrentMaturityReturns(currentMaturity)
	fakeReceipt.TargetMaturityReturns(targetMaturity)

	return fakeReceipt
}
