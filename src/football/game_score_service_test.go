package football_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-backend-test-2/src/football"
)

var _ = Describe("GameScoreService", func() {
	gss := football.NewGameScoreService()

	DescribeTable("Get possible combinations count for game score",
		func (score string, expectedCombinations int) {
			combinations := gss.GetCombinationsCount(score)
			Expect(combinations).To(Equal(expectedCombinations))
		},
		Entry("Valid case 1", "3x15", 4),
	)
})
