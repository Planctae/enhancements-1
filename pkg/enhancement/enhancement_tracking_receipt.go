package enhancement

import (
	yaml "gopkg.in/yaml.v2"
)

type TrackingReceipt interface {
	Title() string
	Authors() []string
	SponsoringSIG() string
	AffectedSubprojects() []string
	KEPLocation() string
	TestLocations() []string
	ReleaseNoteLocation() string
	DocumentationLocations() []string
	ContactEmail() string
	CurrentMaturity() string
	TargetMaturity() string
}

func NewTrackingReceipt(rawBytes []byte) (TrackingReceipt, error) {
	r := &trackingReceipt{}

	err := yaml.Unmarshal(rawBytes, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

type trackingReceipt struct {
	TitleField                  string   `yaml:"title,omitempty"`
	AuthorsField                []string `yaml:"authors,omitempty"`
	SponsoringSIGField          string   `yaml:"sponsoring_sig,omitempty"`
	AffectedSubprojectsField    []string `yaml:"affected_subprojects,omitempty"`
	KEPLocationField            string   `yaml:"kep_location,omitempty"`
	TestLocationsField          []string `yaml:"test_locations,omitempty"`
	ReleaseNoteLocationField    string   `yaml:"release_note_location,omitempty"`
	DocumentationLocationsField []string `yaml:"documentation_locations,omitempty"`
	ContactEmailField           string   `yaml:"contact_email,omitempty"`
	CurrentMaturityField        string   `yaml:"current_maturity,omitempty"`
	TargetMaturityField         string   `yaml:"target_maturity,omitempty"`
}

func (r *trackingReceipt) Title() string                    { return r.TitleField }
func (r *trackingReceipt) Authors() []string                { return r.AuthorsField }
func (r *trackingReceipt) SponsoringSIG() string            { return r.SponsoringSIGField }
func (r *trackingReceipt) AffectedSubprojects() []string    { return r.AffectedSubprojectsField }
func (r *trackingReceipt) KEPLocation() string              { return r.KEPLocationField }
func (r *trackingReceipt) TestLocations() []string          { return r.TestLocationsField }
func (r *trackingReceipt) ReleaseNoteLocation() string      { return r.ReleaseNoteLocationField }
func (r *trackingReceipt) DocumentationLocations() []string { return r.DocumentationLocationsField }
func (r *trackingReceipt) ContactEmail() string             { return r.ContactEmailField }
func (r *trackingReceipt) CurrentMaturity() string          { return r.CurrentMaturityField }
func (r *trackingReceipt) TargetMaturity() string           { return r.TargetMaturityField }
