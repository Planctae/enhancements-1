package render

import (
	"bytes"
	"text/template"

	"github.com/planctae/enhancements-tracking-ng/pkg/render/unrendered"
)

func ReleaseOwners(release string) ([]byte, error) {
	ownersContent := &bytes.Buffer{}

	t, err := template.New("release owners file").Parse(unrendered.ReleaseOwners)
	if err != nil {
		return nil, err
	}

	err = t.Execute(ownersContent, release)
	if err != nil {
		return nil, err
	}

	return ownersContent.Bytes(), nil
}
