package cw_logAnalyzer

import (
	cw_logParser "cloudwalk-assessment/cw-logParser"
	"cloudwalk-assessment/quake3"
	"fmt"
	"strconv"

	funk "github.com/thoas/go-funk"
)

const worldKillerId int = 1022

type GameInfo struct {
	TotalKills int
	Players []string
	Kills map[string]int
}

type MODGameInfo struct {
	KillsByMeans map[string]int
}

func GetGamesInfo(games []cw_logParser.Game) map[string]GameInfo {
	gamesInfo := make(map[string]GameInfo, 0)
	for index, game := range games {
		
		game := GameInfo {
			Players: getPlayerNames(game.Players),
			TotalKills: getTotalKills(game.Kills),
			Kills: getKillsByPlayer(game),
		}
		index := getGameNumber(index)
		fmt.Println(index)
		gamesInfo[index] = game
	}

	return gamesInfo;
}

func GetMODGamesInfo(games []cw_logParser.Game) map[string]MODGameInfo {
	modGamesInfo := make(map[string]MODGameInfo)
	for index, game := range games {
		modGame := MODGameInfo {
			KillsByMeans: getKillsByMean(game.Kills),
		}
		modGamesInfo[getGameNumber(index)] = modGame
	}	
	return modGamesInfo
}

func getGameNumber(gameNumber int) string {
	return "game_"+ strconv.Itoa(gameNumber + 1)
}

func getKillsByMean(kills []cw_logParser.Kill) map[string]int {
	killsByMeant := make(map[string]int)

	for _, kill := range kills {
		meanOfDeathName := quake3.MeanOfDeath(kill.MeanOfDeath).String()
		
		if entry, ok := killsByMeant[meanOfDeathName]; ok {
			entry++
			killsByMeant[meanOfDeathName] = entry
		} else {
			killsByMeant[meanOfDeathName] = 1
		}
	}

	return killsByMeant;
}

func getPlayerNames(gamePlayers map[int]cw_logParser.Player) []string {
		players := funk.Values(gamePlayers).([]cw_logParser.Player)
		playerNames := funk.Get(players, "Name").([]string)
		return playerNames;
}

func getTotalKills(gameKills []cw_logParser.Kill) int{
	return len(gameKills)
}

func getKillsByPlayer(game cw_logParser.Game) map[string]int {
	killsByPlayer := make(map[string]int, 0)

	for _, player := range game.Players {
		killsByPlayer[player.Name] = 0
	}

	for _, kill := range game.Kills {
		playerName := game.Players[kill.KillerId].Name

		if value, ok := killsByPlayer[playerName]; ok{
			value++;
			killsByPlayer[playerName] = value
		} else if kill.KillerId == worldKillerId {
			killedPlayer := game.Players[kill.KilledId].Name
			killScore := killsByPlayer[killedPlayer]
			killScore--
			killsByPlayer[killedPlayer] = killScore
		}
	}

	return killsByPlayer
}