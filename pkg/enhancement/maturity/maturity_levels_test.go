package maturity_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement/maturity"
)

var _ = Describe("Working with enhancement maturity levels", func() {
	Describe("LevelMatches()", func() {
		It("performs a case insenstive comparison between given and desired maturity levels", func() {
			Expect(maturity.LevelMatches("RC", maturity.ReleaseCandidate)).To(BeTrue(), "expected `RC` to match maturity.ReleaseCandidate")
			Expect(maturity.LevelMatches("Rc", maturity.ReleaseCandidate)).To(BeTrue(), "expected `Rc` to match maturity.ReleaseCandidate")
			Expect(maturity.LevelMatches("rc", maturity.ReleaseCandidate)).To(BeTrue(), "expected `rc` to match maturity.ReleaseCandidate")

			Expect(maturity.LevelMatches("ALPHA", maturity.Alpha)).To(BeTrue(), "expected `ALPHA` to match maturity.Alpha")
			Expect(maturity.LevelMatches("Alpha", maturity.Alpha)).To(BeTrue(), "expected `Alpha` to match maturity.Alpha")
			Expect(maturity.LevelMatches("alpha", maturity.Alpha)).To(BeTrue(), "expected `alpha` to match maturity.Alpha")

			Expect(maturity.LevelMatches("BETA", maturity.Beta)).To(BeTrue(), "expected `BETA` to match maturity.Beta")
			Expect(maturity.LevelMatches("Beta", maturity.Beta)).To(BeTrue(), "expected `Beta` to match maturity.Beta")
			Expect(maturity.LevelMatches("beta", maturity.Beta)).To(BeTrue(), "expected `beta` to match maturity.Beta")

			Expect(maturity.LevelMatches("GA", maturity.GenerallyAvailable)).To(BeTrue(), "expected `GA` to match maturity.GenerallyAvailable")
			Expect(maturity.LevelMatches("Ga", maturity.GenerallyAvailable)).To(BeTrue(), "expected `Ga` to match maturity.GenerallyAvailable")
			Expect(maturity.LevelMatches("ga", maturity.GenerallyAvailable)).To(BeTrue(), "expected `ga` to match maturity.GenerallyAvailable")

		})
	})
})
