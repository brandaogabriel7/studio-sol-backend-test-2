package football

type TeamPointsService struct {}

func NewTeamPointsService() *TeamPointsService {
	return &TeamPointsService{}
}

func (tps *TeamPointsService) GetTeamCombinationsCount(teamPoints int) int {
	if teamPoints < 0 {
		return 0
	}
	combinations := make([]int, teamPoints + 1)
	combinations[0] = 1

	// field goal, touchdown, touchdown + 1 extra, touchdown + 2 extra
	possiblePoints := []int { 3, 6, 7, 8 }

	for j := 0; j < len(possiblePoints); j++ {
		for i := possiblePoints[j]; i < len(combinations); i++ {
			combinations[i] += combinations[i - possiblePoints[j]]
		}
	}

	return combinations[teamPoints]
}