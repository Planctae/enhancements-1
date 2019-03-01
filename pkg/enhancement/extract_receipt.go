package enhancement

import (
	"bytes"
	"fmt"
)

func ExtractFromMarkdown(input []byte) (TrackingReceipt, error) {
	openingLoc := bytes.Index(input, []byte(codeblockSep))
	closingLoc := bytes.LastIndex(input, []byte(codeblockSep))

	if openingLoc == closingLoc {
		return nil, fmt.Errorf("could not determine receipt location: could not find opening and closing code block separators\n input: \n%s", string(input))
	}

	receiptBytes := input[openingLoc + len(codeblockSep):closingLoc]

	r, err := NewTrackingReceipt(receiptBytes)
	if err != nil {
		return nil, err
	}

	return r, nil
}

const (
	codeblockSep = "```"
)
