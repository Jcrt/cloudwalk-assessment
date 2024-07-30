// Package cw_logAnalyzer group all required artifacts to analyze the logs generated by log parser
// and aggregate all relevant information from game log structs
package cw_logAnalyzer

import (
	cw_logParser "cloudwalk-assessment/cw-logParser"
	"cloudwalk-assessment/quake3"
	"strconv"

	orderedMap "github.com/iancoleman/orderedmap"
	funk "github.com/thoas/go-funk"
)

// Defines the <world> code in the Kill logs
const worldKillerId int = 1022

//Interface ILogAnalyzer defines all methods offered by log analyzer
type ILogAnalyzer interface {
	GetMatchesInfo(parsedLog cw_logParser.LogParser) orderedMap.OrderedMap;
	GetMODMatchesInfo(parsedLog cw_logParser.LogParser) orderedMap.OrderedMap;
}

// Struct LogAnalyzer is used as ILogAnalyzer interface implementation
type LogAnalyzer struct {

}

// MatchInfo struct defines all information relative to the Game in analysis aspect
type MatchInfo struct {
	TotalKills int
	Players []string
	Kills orderedMap.OrderedMap
	Ranking orderedMap.OrderedMap
}

// MODMatchInfo struct defines a map of all deaths occurred during the game and group it providing a counter of each Mean of Death
type MODMatchInfo struct {
	KillsByMeans map[string]int
}

// Func GetMatchesInfo build a map containing all matches and it's relevant information
// The parameter games receives an array of game log parser object
//
// Returns an OrderedMap containing all matches
func (logAnalyzer LogAnalyzer) GetMatchesInfo(parsedLog cw_logParser.LogParser) orderedMap.OrderedMap {
	matchesInfo := orderedMap.New()

	for index, match := range parsedLog.Matches {
		kills := getKillsByPlayer(match)
		currentMatch := MatchInfo {
			Players: getPlayerNames(match.Players),
			TotalKills: getTotalKills(match.Kills),
			Kills: kills,
			Ranking: getRankingByKills(kills),
		}
		index := getMatchIdentifier(index)
		matchesInfo.Set(index, currentMatch)
	}

	return *matchesInfo;
}

// Func GetMODMatchesInfo build a map for each mean of death that occurred in logs, group it and count it
// The parameter games receives an array of game log parser object
//
// Returns an OrderedMap containing all means of death grouped and counted
func (logAnalyzer LogAnalyzer) GetMODMatchesInfo(parsedLog cw_logParser.LogParser) orderedMap.OrderedMap {
	modMatchesInfo := orderedMap.New()
	for index, match := range parsedLog.Matches {
		currentMODMatch := MODMatchInfo {
			KillsByMeans: getKillsByMean(match.Kills),
		}
		modMatchesInfo.Set(getMatchIdentifier(index), currentMODMatch)
	}	
	return *modMatchesInfo
}

func getMatchIdentifier(matchNumber int) string {
	return "game_"+ strconv.Itoa(matchNumber + 1)
}

// Func getKillsByMean generates a map containing all means of death that occurred during a match
// The parameter kills receive an array of Kill and group it into a map with a counter
//
// Returns a map containing as key the name of mean of death and as value the amount of time that it occurs
func getKillsByMean(kills []cw_logParser.Kill) map[string]int {
	killsByMean := make(map[string]int)

	for _, kill := range kills {
		meanOfDeathName := quake3.MeanOfDeath(kill.MeanOfDeath).String()
		
		if entry, ok := killsByMean[meanOfDeathName]; ok {
			entry++
			killsByMean[meanOfDeathName] = entry
		} else {
			killsByMean[meanOfDeathName] = 1
		}
	}

	return killsByMean;
}

// Func getPlayerNames generates a string array containing all player names that have entered during a match
// The parameter gamePlayers receive a map of Player gathered from logs
//
// Returns a string array containing all player names at the end of the match
func getPlayerNames(matchPlayers map[int]cw_logParser.Player) []string {
		players := funk.Values(matchPlayers).([]cw_logParser.Player)
		playerNames := funk.Get(players, "Name").([]string)
		return playerNames;
}

// Func getTotalKills counts all kills occurred in a match
// The parameter gameKills receives a list of Kill and count it
//
// Returns an int with the total amount of kills occurred in the match
func getTotalKills(matchKills []cw_logParser.Kill) int{
	return len(matchKills)
}

// Func getKillsByPlayer groups all kills by player and count it
// If a player was killed by the <world>, it's scored is decreased by 1 every time that it happened
// The parameter game is the Game struct
//
// Returns an OrderedMap where the key is the mean of kill and the value is an int as a counter of occurrences
func getKillsByPlayer(match cw_logParser.Match) orderedMap.OrderedMap {
	killsByPlayer := orderedMap.New()

	for _, player := range match.Players {
		killsByPlayer.Set(player.Name, 0)
	}

	for _, kill := range match.Kills {
		playerName := match.Players[kill.KillerId].Name

		if value, ok := killsByPlayer.Get(playerName); ok{
			intValue := value.(int)
			intValue++;
			killsByPlayer.Set(playerName, intValue)
		} else if kill.KillerId == worldKillerId {
			killedPlayer := match.Players[kill.KilledId].Name
			killScore, _ := killsByPlayer.Get(killedPlayer)
			intKillScore := killScore.(int)
			intKillScore--
			killsByPlayer.Set(killedPlayer, intKillScore)
		}
	}

	return *killsByPlayer
}

// Func getRankingByKills creates a copy of kills orderedMap and with it's data creates a new ranking list
// based on the kills list
// The parameter kills is an orderedMap containing players and amount of kills unordered
// Returns a new orderedList containing the ranking

// OBS: In terms of performance, it's not good because I'm recreating a list and doubling the memory consumption
// I only did it to clearly separate the concerns of the functions, but I could do this in one function that would
// return both kills ordered and ranking 
func getRankingByKills(kills orderedMap.OrderedMap) orderedMap.OrderedMap {

	rankedPlayers := orderedMap.New()
	killCopy := orderedMap.New()

	for _, kill := range kills.Keys() {
		value, _ := kills.Get(kill)
		killCopy.Set(kill, value)
	}

	killCopy.Sort(func(a *orderedMap.Pair, b *orderedMap.Pair) bool{
		return a.Value().(int) > b.Value().(int)
	})

	for index, rankedPlayer := range killCopy.Keys() {
		rankedPlayers.Set(strconv.Itoa(index + 1), rankedPlayer)
	}

	return *rankedPlayers
}

