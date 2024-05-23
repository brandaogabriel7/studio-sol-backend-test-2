package integration_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-backend-test-2/graph"
	"github.com/brandaogabriel7/studio-sol-backend-test-2/graph/model"
	"github.com/brandaogabriel7/studio-sol-backend-test-2/src/football"
)

var _ = Describe("Verify", func() {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			GameScoreService: football.NewGameScoreService(),
		},
	})))

	Context("Calculating number of possible combinations to get score", func() {
		DescribeTable("score, combinations",
			func(score string, expectedCombinations int) {
				var resp struct {
					Verify model.Verify
				}

				error := c.Post(
					`
					mutation($score: String!) {
						verify(score: $score) {
							combinations
						}
					}
					`,
					&resp,
					client.Var("score", score),
				)

				Expect(error).NotTo(HaveOccurred())
				Expect(resp.Verify.Combinations).To(Equal(expectedCombinations))
			},
			Entry("Test case 1", "3x15", 4),
			Entry("Test case 2", "8x5", 0),
		)
	})
})
