package football_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-backend-test-2/src/football"
)

var _ = Describe("TeamPointsService", func() {

	tps := football.NewTeamPointsService()

	DescribeTable("Get the number of possible combinations for the given team points in a football game",
		func (teamPoints int, expectedCombinations int) {
			combinations := tps.GetTeamCombinationsCount(teamPoints)

			Expect(combinations).To(Equal(expectedCombinations))
		},
		Entry("Valid test case 1", 3, 1),
		Entry("Valid test case 2", 0, 1),
	)

	DescribeTable("Return 0 for team points that are not possible to get in football",
		func (teamPoints int) {
			combinations := tps.GetTeamCombinationsCount(teamPoints)
			
			Expect(combinations).To(Equal(0))
		},
		Entry("Invalid test case 1", 2),
		Entry("Invalid test case 2", 4),
	)
})
