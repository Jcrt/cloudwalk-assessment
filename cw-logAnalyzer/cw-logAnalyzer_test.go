package cw_logAnalyzer

import (
	cw_logParser "cloudwalk-assessment/cw-logParser"
	"cloudwalk-assessment/quake3"
	"testing"

	"github.com/stretchr/testify/assert"
)

var logAnalizer ILogAnalyzer = &LogAnalyzer{}

func Test_GetMatchesInfo_ShouldReturnCorrectReport(t *testing.T) {
	var log cw_logParser.Log

	output := logAnalizer.GetMatchesInfo(log)

	assert.NotNil(t, output)
}

func Test_GetMatchesInfo_ShouldReturnCorrectPlayerNames(t *testing.T) {
	
	players, playersMap := getChoqueDeCulturaPlayersMap()

	matches := []cw_logParser.Match{
		{
			Order: 1,
			Players: playersMap,
		},
	}

	log := cw_logParser.Log{
		Matches: matches,
	}

	output := logAnalizer.GetMatchesInfo(log)

	assert.NotNil(t, output)

	match1, _ := output.Get("game_1")
	castedMatch := match1.(MatchInfo)

	assert.Contains(t, castedMatch.Players, players[0].Name)
	assert.Contains(t, castedMatch.Players, players[1].Name)
	assert.Contains(t, castedMatch.Players, players[2].Name)
	assert.Contains(t, castedMatch.Players, players[3].Name)
}

func Test_GetMatchesInfo_ShouldReturnCorrectKills(t *testing.T) {
	
	_, playersMap := getChoqueDeCulturaPlayersMap()

	testCases := []struct {
		playerIndex, amountOfKills int
	}{
		{ 1, 2 }, 
		{ 2, 1 },
		{ 3, 1 },
		{ 4, 0 },
	}

	matches := []cw_logParser.Match{
		{
			Order: 1,
			Players: playersMap,
			Kills: getKills(),
		},
	}

	log := cw_logParser.Log{
		Matches: matches,
	}

	output := logAnalizer.GetMatchesInfo(log)
	assert.NotNil(t, output)

	match1, _ := output.Get("game_1")
	castedMatch := match1.(MatchInfo)
	
	assert.Equal(t, castedMatch.TotalKills, len(getKills()))

	for _, data := range testCases {
		playerKill, exists:= castedMatch.Kills.Get(playersMap[data.playerIndex].Name)
		castedPlayer := playerKill.(int)
		assert.True(t, exists)
		assert.Equal(t, castedPlayer, data.amountOfKills)
	}
}

func Test_GetMatchesInfo_ShouldReturnCorrectRanking(t *testing.T) {
	
	players, playersMap := getChoqueDeCulturaPlayersMap()

	testCases := []struct {
		playerIndex string
		playerName string
	}{
		{ "1", players[0].Name }, 
		{ "2", players[1].Name },
		{ "3", players[2].Name },
		{ "4", players[3].Name },
	}

	matches := []cw_logParser.Match{
		{
			Order: 1,
			Players: playersMap,
			Kills: getKills(),
		},
	}

	log := cw_logParser.Log{
		Matches: matches,
	}

	output := logAnalizer.GetMatchesInfo(log)
	assert.NotNil(t, output)

	match1, _ := output.Get("game_1")
	castedMatch := match1.(MatchInfo)
	
	assert.Equal(t, castedMatch.TotalKills, len(getKills()))

	for _, data := range testCases {
		playerName, exists:= castedMatch.Ranking.Get(data.playerIndex)
		castedPlayer := playerName.(string)
		assert.True(t, exists)
		assert.Equal(t, castedPlayer, data.playerName)
	}
}

func Test_GetMODMatchesInfo_ShouldReturnCorrectValues(t *testing.T){
	_, playersMap := getChoqueDeCulturaPlayersMap()

	testCases := []struct {
		MODIndex string
		AmountOfMOD int
	}{
		{ quake3.MOD_MACHINEGUN.String(), 1 }, 
		{ quake3.MOD_RAILGUN.String(), 1 },
		{ quake3.MOD_ROCKET.String(), 1 },
		{ quake3.MOD_CRUSH.String(), 1 },
		{ quake3.MOD_CHAINGUN.String(), 2 },
	}

	matches := []cw_logParser.Match{
		{
			Order: 1,
			Players: playersMap,
			Kills: getKills(),
		},
	}

	log := cw_logParser.Log{
		Matches: matches,
	}

	output := logAnalizer.GetMODMatchesInfo(log)
	assert.NotNil(t, output)

	match1, _ := output.Get("game_1")
	castedMatch := match1.(MODMatchInfo)

	for _, data := range testCases {
		amountOfMOD, exists:= castedMatch.KillsByMeans[data.MODIndex]
		assert.True(t, exists)
		assert.Equal(t, data.AmountOfMOD, amountOfMOD)
	}
}

func getChoqueDeCulturaPlayersMap() ([]cw_logParser.Player, map[int]cw_logParser.Player) {
	players := []cw_logParser.Player {
			{ Name: "Rogerinho do Ingá" }, 
			{ Name: "Julinho da Van" }, 
			{ Name: "Renan" },
			{ Name: "Maurílio" },
		}

	playersMap := make(map[int]cw_logParser.Player)
	playersMap[1] = players[0]
	playersMap[2] = players[1]
	playersMap[3] = players[2]
	playersMap[4] = players[3]

	return players, playersMap
}

func getKills() []cw_logParser.Kill {
	kills := []cw_logParser.Kill {
		{ KillerId: 1, KilledId: 2, MeanOfDeath: quake3.MOD_MACHINEGUN }, 
		{ KillerId: 2, KilledId: 3, MeanOfDeath: quake3.MOD_RAILGUN }, 
		{ KillerId: 3, KilledId: 4, MeanOfDeath: quake3.MOD_ROCKET }, 
		{ KillerId: 4, KilledId: 1, MeanOfDeath: quake3.MOD_CRUSH }, 
		{ KillerId: 1, KilledId: 4, MeanOfDeath: quake3.MOD_CHAINGUN }, 	
		{ KillerId: 1022, KilledId: 4, MeanOfDeath: quake3.MOD_CHAINGUN },
	}

	return kills
}