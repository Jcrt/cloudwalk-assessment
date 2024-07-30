// Package cw_logParser is the package responsible by parse all log string information to solid types
// offering useful structures to achieve this goal
package cw_logParser

import (
	cw_error "cloudwalk-assessment/cw-error"
	quake3 "cloudwalk-assessment/quake3"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const prefixRegexp string = `(?i)(?<prefix>\d{1,2}:\d{2} )`

// keywordRegex will be instantiated only here to improve performance
var keywordRegex = regexp.MustCompile(prefixRegexp + `(?<keyword>[^:]*:)`)
var killKeywordRegex = regexp.MustCompile(prefixRegexp + `(?<keyword>`+ string(SK_Kill) + ` )(?<data>((\d{1,4} ?){3}))`)
var playerKeywordRegex = regexp.MustCompile(prefixRegexp + `(?<keyword>`+ string(SK_ClientUserinfoChanged) + ` )(?<info>\d* n\\[^\\t]*)`)

//Interface ILogParser defines all methods offered by log parser
type ILogParser interface {
	Parse(logString string) (Log, error);
}

// Struct LogParser is used as ILogParser interface implementation
type LogParser struct {
	errorHandler cw_error.IErrorHandler
}

func CreateLogParser(errorHandler cw_error.IErrorHandler) *LogParser {
	return &LogParser{
		errorHandler: errorHandler,
	}
}

// Struct Log is used as ILogParser interface implementation
type Log struct {
	Matches []Match;
}

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
	SK_InitGame searchingKeywords = "initgame:"
	SK_ClientUserinfoChanged searchingKeywords = "clientuserinfochanged:"
	SK_Kill searchingKeywords = "kill:"
)

// ParseLog is the function that retrieve log info and parse it to log parser types
func (logParser *LogParser) Parse(logString string) (Log, error) {
	matches := make([]Match, 0)
	trimmedLogString := strings.TrimSpace(logString)
	lines := strings.SplitN(trimmedLogString, "\n", -1)
	errors := make([]error, 0)

	isValid, err := validate(logString)

	if(isValid) {
		for index, line := range lines{
			hasKeyword, keyword := findKeyword(line)
			if(hasKeyword){
					switch op := keyword; strings.ToLower(op) {
						case string(SK_InitGame): {
							currentMatch := Match{
								Order: len(matches) + 1,
								Players: make(map[int]Player, 0),
								Kills: make([]Kill, 0),
							}
							matches = append(matches, currentMatch)
						}
						case string(SK_ClientUserinfoChanged): {
							if(len(matches) > 0){
								exists, player, playerId, playerErr := logParser.parsePlayerLine(line)

								if(exists) {
									matches[len(matches)-1].Players[playerId] = player
								} else {
									errors = append(errors, fmt.Errorf("line %s: %s", strconv.Itoa(index), playerErr.Error()))
								}
							}
						}
						case string(SK_Kill): {
							if(len(matches) > 0){
								exists, kill, killErr := logParser.parseKillLine(line)

								if(exists) {
									matches[len(matches)-1].Kills = append(matches[len(matches)-1].Kills, kill)
								} else {
									errors = append(errors, fmt.Errorf("line %s: %s", strconv.Itoa(index), killErr.Error()))
								}
							}
						}
						default: continue; 
					}
			}
		}
	} else {

		err = logParser.errorHandler.BuildErrorOutput(errors...)
	}

	parsedLog := Log {
		Matches: matches,
	}

	return  parsedLog, err
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

func findSubgroupRegex(line string, regexp *regexp.Regexp, matchIndex int) (bool, string) {
	match := regexp.FindStringSubmatch(line)
	contains := false
	keyword := ""

	if(len(match) >= matchIndex){
		contains = true
		keyword = match[matchIndex]
	} 

	return contains, keyword
}

// Func parsePlayerLine is used to parse a log line that defines a player to the Player struct
// The parameter line is a log line
//
// Returns a tuple where the first item is the Player struct and the second is an int containing player id
func (logParser *LogParser) parsePlayerLine(line string) (bool, Player, int, error) {

	var player Player
	var playerId int
	var err error
	
	exists, data := findSubgroupRegex(line, playerKeywordRegex, 3)
	before, after, _ := strings.Cut(data, "n\\")


	if(exists) {
		playerId, _ = strconv.Atoi(strings.TrimSpace(before))

		player = Player {
			Name: after,
		}
	} else{
		err = fmt.Errorf("not well formed player log line: %q", line)
	}

	return exists, player, playerId, err;	
}

// Func parseKillLine is used to parse a log line that defines a kill to the Kill struct
// The parameter line is a log line
//
// Returns a Kill struct containing all required kill data
func (logParser *LogParser) parseKillLine(line string) (bool, Kill, error) {
	var kill Kill
	var err error

	exists, data := findSubgroupRegex(line, killKeywordRegex, 3)

	if(!exists){
		err = fmt.Errorf("not well formed Kill log line: %q", line)
	} else {
		infos := strings.Split(strings.TrimSpace(data), " ")
		killerId, killerErr := strconv.Atoi(infos[0])
		killedId, killedErr := strconv.Atoi(infos[1])
		meanOfDeath, modErr := strconv.Atoi(infos[2])

		err = logParser.errorHandler.BuildErrorOutput(killerErr, killedErr, modErr)

		if(err == nil) {
			kill = Kill{
				KillerId: killerId,
				MeanOfDeath: quake3.MeanOfDeath(meanOfDeath),
				KilledId: killedId,
			}
		}
	}

	return exists, kill, err;
}