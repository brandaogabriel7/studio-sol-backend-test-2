package sum_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSum(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sum Suite")
}
