// Package cw_logParser is the package responsible by parse all log string information to solid types
// offering useful structures to achieve this goal
package cw_logParser

import (
	cw_logreader "cloudwalk-assessment/cw-logReader"
	quake3 "cloudwalk-assessment/quake3"
	"strconv"
	"strings"
)

// Game is the struct that defines the useful information about each match in the log
type Game struct {
	Order int
	Players map[int]Player
	Kills []Kill
}

// Player is the struct that defines the useful information about the player during a match
type Player struct {
	Name string 
}

// Kill is the struct that defines the useful information about each kill log entry that occurs during a match
type Kill struct {
	KillerId int
	KilledId int
	MeanOfDeath quake3.MeanOfDeath
}

// searchingKeywords defines the required keywords that should be found in the match log and that will scope
// the log parser objects
type searchingKeywords string
const(
	SK_InitGame searchingKeywords = "InitGame:"
	SK_ClientUserinfoChanged searchingKeywords = "ClientUserinfoChanged:"
	SK_Kill searchingKeywords = "Kill:"
)

// ParseLog is the function that retrieve log info and parse it to log parser types
func ParseLog() []Game {
	games := make([]Game, 0)
	fileContent := cw_logreader.ReadLog()
	lines := strings.SplitN(fileContent, "\n", -1)

	for _, line := range lines{
		//TODO: Use some better way to compare keywords as switch extracting keywords with regex maybe
		if(searchingKeywordExists(line, SK_InitGame)){
			currentGame := Game{
				Order: len(games) + 1,
				Players: make(map[int]Player, 0),
			}
			games = append(games, currentGame)
		} else if(searchingKeywordExists(line, SK_ClientUserinfoChanged)){
			player, playerId := parsePlayerLine(line)
			games[len(games)-1].Players[playerId] = player
		} else if(searchingKeywordExists(line, SK_Kill)) {
			kill := parseKillLine(line)
			games[len(games)-1].Kills = append(games[len(games)-1].Kills, kill)
		}
	}

	return games
}

// Func searchingKeywrodExists searches inside a log line if words of interest whether present or not
// The parameter line receives a log line
// The parameter searchingKeyword is the word of interest that's mapped in searchingKeywords enum
//
// Returns true if word of interest was found, otherwise false
func searchingKeywordExists(line string, searchingKeyword searchingKeywords) bool {
	contains := strings.Contains(line, string(searchingKeyword))
	return contains
}

// Func parsePlayerLine is used to parse a log line that defines a player to the Player struct
// The parameter line is a log line
//
// Returns a tuple where the first item is the Player struct and the second is an int containing player id
func parsePlayerLine(line string) (Player, int) {
	before, _, _ := strings.Cut(line, "\\t")
	before, playerName, _ := strings.Cut(before,"n\\")
	_, playerId, _ := strings.Cut(before, string(SK_ClientUserinfoChanged))
	playerIdInt, _ := strconv.Atoi(strings.TrimSpace(playerId))

	player := Player {
		Name: playerName,
	}

	return player, playerIdInt;	
}

// Func parseKillLine is used to parse a log line that defines a kill to the Kill struct
// The parameter line is a log line
//
// Returns a Kill struct containing all required kill data
func parseKillLine(line string) Kill {
	_, after, _ := strings.Cut(line, string(SK_Kill))
	before, _, _ := strings.Cut(after, ":")
	infos := strings.Split(strings.TrimSpace(before), " ")

	killerId, _ := strconv.Atoi(infos[0])
	killedId, _ := strconv.Atoi(infos[1])
	meanOfDeath, _ := strconv.Atoi(infos[2])

	kill := Kill{
		KillerId: killerId,
		MeanOfDeath: quake3.MeanOfDeath(meanOfDeath),
		KilledId: killedId,
	}

	return kill;
}