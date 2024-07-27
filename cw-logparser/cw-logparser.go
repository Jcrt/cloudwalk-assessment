// Package cw_logParser is the package responsible by parse all log string information to solid types
// offering useful structures to achieve this goal
package cw_logParser

import (
	quake3 "cloudwalk-assessment/quake3"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// keywordRegex will be instantiated only here to improve performance
var keywordRegex = regexp.MustCompile(`(?<prefix>\d{1,2}:\d{2} )(?<keyword>[^:]*:)`)

// Match is the struct that defines the useful information about each match in the log
type Match struct {
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
func ParseLog(logString string) ([]Match, error) {
	matches := make([]Match, 0)
	lines := strings.SplitN(logString, "\n", -1)

	isValid, err := validate(logString)

	if(isValid) {
		for _, line := range lines{
			hasKeyword, keyword := findKeyword(line)
			if(hasKeyword){
					switch op := keyword; op {
						case string(SK_InitGame): {
							currentMatch := Match{
								Order: len(matches) + 1,
								Players: make(map[int]Player, 0),
							}
							matches = append(matches, currentMatch)
						}
						case string(SK_ClientUserinfoChanged): {
							if(len(matches) > 0){
								player, playerId := parsePlayerLine(line)
								matches[len(matches)-1].Players[playerId] = player
							}
						}
						case string(SK_Kill): {
							if(len(matches) > 0){
								kill := parseKillLine(line)
								matches[len(matches)-1].Kills = append(matches[len(matches)-1].Kills, kill)
							}
						}
						default: continue; 
					}
			}
		}
	} else {
		err = fmt.Errorf("the parse process failed. reason: %s", err.Error())
		log.Output(1, err.Error())
	}

	return matches, err
}

//Func validate makes a basic input validation
// The parameter logString receives the string containing log
//
// Returns a bool sinalizing if input is valid or not and error variable
func validate(logString string) (bool, error) {
	var err error

	if(len(logString) == 0){
		err = fmt.Errorf("log content is empty")
		return false, err
	}
		
	return true, err
}

// Func findKeyword searches inside a log line if words of interest whether present or not
// Regexp here wouldn't be the most efficient process to find keywords, but for this assessment 
// I've decided that should be interesting show the usage of regex 😃
// The parameter line receives a log line
//
// Returns true if word of interest was found and the keyword, otherwise false and empty string
func findKeyword(line string) (bool, string) {
	match := keywordRegex.FindStringSubmatch(line)
	contains := false
	keyword := ""
	if(len(match) > 1 && keywordRegex.SubexpIndex("keyword") > -1){
		contains = true
		keyword = match[keywordRegex.SubexpIndex("keyword")]
	} 
	return contains, keyword
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