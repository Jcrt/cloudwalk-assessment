// Package cw_reporter is responsible by provide all available reports over game log to final user
// This package exposes all report types and interact with all other packages in order to encapsulate other application layers
package cw_reporter

import (
	cw_logAgnalyzer "cloudwalk-assessment/cw-logAnalyzer"
	cw_logParser "cloudwalk-assessment/cw-logParser"
	cw_logReader "cloudwalk-assessment/cw-logReader"
)

// Func GetGameReport generates the basic game report as requested in CloudWalk Assessment item 3.2
//
// Returns a map where the key is the match number and the value is the match info
func GetGameReport() map[string]cw_logAgnalyzer.GameInfo{
	stringLog := cw_logReader.GetLog()
	games := cw_logParser.ParseLog(stringLog)
	analyzedGames := cw_logAgnalyzer.GetGamesInfo(games)
	return analyzedGames
}

// Func GetGameReport generates the MOD game report as requested in CloudWalk Assessment item 3.3
//
// Returns a map where the key is the MOD type and the value is the count of times that happened
func GetMODGameReport() map[string]cw_logAgnalyzer.MODGameInfo{
	stringLog := cw_logReader.GetLog()
	games := cw_logParser.ParseLog(stringLog)
	analyzedGames := cw_logAgnalyzer.GetMODGamesInfo(games)
	return analyzedGames
}