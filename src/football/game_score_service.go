package football

import (
	"strconv"
	"strings"
)

type GameScoreService struct {
	tps *TeamPointsService
}

func NewGameScoreService() *GameScoreService {
	tps := NewTeamPointsService()

	
	return &GameScoreService{ tps: tps }
}

func (gss *GameScoreService) GetCombinationsCount (score string) int {
	team1, team2 := getTeamsPointsFromScore(score)
	
	return gss.tps.GetTeamCombinationsCount(team1) * gss.tps.GetTeamCombinationsCount(team2)
}

func getTeamsPointsFromScore(score string) (int, int) {
	const scoreDelimiter string = "x"

	teamScores := strings.Split(score, scoreDelimiter)
	
	team1, err := strconv.ParseInt(teamScores[0], 0, 0)
	if err != nil {
		return 0, 0
	}

	team2, err := strconv.ParseInt(teamScores[1], 0, 0)
	if err != nil {
		return 0, 0
	}

	return int(team1), int(team2)
}