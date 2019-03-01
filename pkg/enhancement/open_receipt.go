package enhancement

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func OpenReceipt(p string) (TrackingReceipt, error) {
	rawBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(p) {
	case markdownExtension:
		return ExtractFromMarkdown(rawBytes)

	case yamlExtension:
		return NewTrackingReceipt(rawBytes)

	default:
		return nil, fmt.Errorf("no rules to handle path without extension: %s", p)

	}
}

const (
	markdownExtension = ".md"
	yamlExtension     = ".yaml"
)
