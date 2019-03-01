package render_test

import (
	"gopkg.in/yaml.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/planctae/enhancements-tracking-ng/pkg/render"
)

var _ = Describe("Rendering an OWNERS file for a release tracking directory", func() {
	Describe("ReleaseOwners()", func() {
		It("renders a simple OWNERS file", func() {
			By("creating an OWNERS file to be used with a centralized OWNERS_ALIASES")

			releaseName := "release-1.15"
			enhancementsTeam := "release-1.15-enhancements-team"
			releaseTeamLeads := "release-1.15-release-team-leads"

			ownersBytes, err := render.ReleaseOwners(releaseName)
			Expect(err).ToNot(HaveOccurred(), "expected no error when rendering a release tracking directory OWNERS file")


			var owners struct {
				Reviewers []string `yaml:"reviewers"`
				Approvers []string `yaml:"approvers"`
			}

			err = yaml.Unmarshal(ownersBytes, &owners)
			Expect(err).ToNot(HaveOccurred(), "expected no error when unmarshalling a rendered OWNERS file as YAML")

			Expect(owners.Reviewers).To(HaveLen(2), "expected 2 reviewers to be listed in a rendered OWNERS file")
			Expect(owners.Approvers).To(HaveLen(2), "expected 2 approvers to be listed in a rendered OWNERS file")

			Expect(owners.Reviewers).To(ConsistOf([]string{enhancementsTeam, releaseTeamLeads}), "expected reviwers to consist of release-1.15-{enhancements-team, release-team-leads}")
			Expect(owners.Approvers).To(ConsistOf([]string{enhancementsTeam, releaseTeamLeads}), "expected reviwers to consist of release-1.15-{enhancements-team, release-team-leads}")
		})
	})
})
