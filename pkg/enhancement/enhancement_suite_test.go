package enhancement_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEnhancement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Enhancement Suite")
}
