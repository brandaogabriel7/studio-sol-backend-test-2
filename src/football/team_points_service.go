package football

type TeamPointsService struct {
	cachedCombinations map[int]int
}

func NewTeamPointsService() *TeamPointsService {
	return &TeamPointsService{
		cachedCombinations: map[int]int{
			0: 1,
		},
	}
}

func (tps *TeamPointsService) GetTeamCombinationsCount(teamPoints int) int {
	if teamPoints < 0 {
		return 0
	}

	if combinations, cached := tps.cachedCombinations[teamPoints]; cached {
		return combinations
	}

	combinations := make([]int, teamPoints+1)
	combinations[0] = 1

	// field goal, touchdown, touchdown + 1 extra, touchdown + 2 extra
	possiblePoints := []int{3, 6, 7, 8}

	for _, points := range possiblePoints {
		for i := points; i < len(combinations); i++ {
			combinations[i] += combinations[i-points]
			tps.cachedCombinations[i] = combinations[i]
		}
	}

	return tps.cachedCombinations[teamPoints]
}
