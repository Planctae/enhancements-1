package maturity_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaturity(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maturity Suite")
}
