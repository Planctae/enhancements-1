package unrendered

const EnhancementsCSV = `
Title,Authors,SIG,Subprojects,KEP,Tests,Release Note,Docs,Contact Email,Current Maturity,Target Maturity
{{- range . }}
{{.Title}},{{joinSpace .Authors}},{{.SponsoringSIG}},{{joinSpace .AffectedSubprojects}},{{.KEPLocation}},{{joinSpace .TestLocations}},{{.ReleaseNoteLocation}},{{joinSpace .DocumentationLocations}},{{.ContactEmail}},{{.CurrentMaturity}},{{.TargetMaturity -}}
{{ end -}}
`

