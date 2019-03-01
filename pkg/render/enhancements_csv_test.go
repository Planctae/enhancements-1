package render_test

import (
	"bytes"
	"encoding/csv"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement"
	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/maturity"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/enhancementfakes"

	"github.com/planctae/enhancements-tracking-ng/pkg/render"
)

var _ = Describe("Rending a list of enhancement tracking receipts to a CSV byte slice", func() {
	Describe("EnhancementsCSV", func() {
		It("produces a CSV byte slice", func() {
			fakeReceiptOne := generateValidReceipt()
			fakeReceiptTwo := generateValidReceipt()
			fakeReceiptThree := generateValidReceipt()
			fakeReceiptFour := generateValidReceipt()

			renderedBytes, err := render.EnhancementsCSV([]enhancement.TrackingReceipt{fakeReceiptOne, fakeReceiptTwo, fakeReceiptThree, fakeReceiptFour})
			Expect(err).ToNot(HaveOccurred(), "expected no error when rendering a CSV byte slice of enhancement tracking receipts")

			csvReader := csv.NewReader(bytes.NewBuffer(renderedBytes))
			_, err = csvReader.ReadAll()
			Expect(err).ToNot(HaveOccurred(), "expected no error when reading a CSV byte slice of enhancement tracking receipts")
		})
	})
})

func generateValidReceipt() *enhancementfakes.FakeTrackingReceipt {
	fakeReceipt := &enhancementfakes.FakeTrackingReceipt{}

	var title string = "ok"
	var authors = []string{"ok"}
	var sponsoringSIG string = "ok"
	var affectedSubprojects = []string{"ok"}
	var kepLocation string = "ok"
	var testLocations = []string{"ok"}
	var releaseNoteLocation string = "ok"
	var documentationLocations = []string{"ok"}
	var contactEmail string = "ok"

	var currentMaturity string = maturity.ReleaseCandidate // must be one of RC|alpha|beta|GA
	var targetMaturity string = maturity.Alpha             // must be one of RC|alpha|beta|GA

	fakeReceipt.TitleReturns(title)
	fakeReceipt.AuthorsReturns(authors)
	fakeReceipt.SponsoringSIGReturns(sponsoringSIG)
	fakeReceipt.AffectedSubprojectsReturns(affectedSubprojects)
	fakeReceipt.KEPLocationReturns(kepLocation)
	fakeReceipt.TestLocationsReturns(testLocations)
	fakeReceipt.ReleaseNoteLocationReturns(releaseNoteLocation)
	fakeReceipt.DocumentationLocationsReturns(documentationLocations)
	fakeReceipt.ContactEmailReturns(contactEmail)
	fakeReceipt.CurrentMaturityReturns(currentMaturity)
	fakeReceipt.TargetMaturityReturns(targetMaturity)

	return fakeReceipt
}
