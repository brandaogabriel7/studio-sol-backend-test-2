package sum_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-backend-test-2/src/sum"
)

var _ = Describe("Sum", func() {
	When("sum is called with 1 and 1", func() {
		It("should return 2", func() {
			Expect(sum.Sum(1, 1)).To(Equal(2))
		})
	})
})
