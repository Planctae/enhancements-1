package render

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/planctae/enhancements-tracking-ng/pkg/enhancement"
	"github.com/planctae/enhancements-tracking-ng/pkg/render/unrendered"
)

func EnhancementsCSV(receipts []enhancement.TrackingReceipt) ([]byte, error) {
	csvContent := &bytes.Buffer{}

	funcMap := template.FuncMap{
		"joinSpace": joinSpace,
	}

	t, err := template.New("enhancements spreadsheet").Funcs(funcMap).Parse(unrendered.EnhancementsCSV)
	if err != nil {
		return nil, err
	}

	err = t.Execute(csvContent, receipts)
	if err != nil {
		return nil, err
	}

	return csvContent.Bytes(), nil
}

func joinSpace(ss []string) string {
	return strings.Join(ss, " ")
}
