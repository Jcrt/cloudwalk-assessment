// Package cw_reporter is responsible by provide all available reports over game log to final user
// This package exposes all report types and interact with all other packages in order to encapsulate other application layers
package cw_reporter

import (
	cw_error "cloudwalk-assessment/cw-error"
	cw_logAgnalyzer "cloudwalk-assessment/cw-logAnalyzer"
	cw_logParser "cloudwalk-assessment/cw-logParser"
	cw_logReader "cloudwalk-assessment/cw-logReader"

	orderedMap "github.com/iancoleman/orderedmap"
)

type ReportType string
const (
	GAME_REPORT ReportType = "game"
	MOD_REPORT ReportType = "mod"
)

//Interface ILogReporter defines all methods offered by log reporter
type ILogReporter interface {
	GetGameReport(reportType ReportType) (IGameReport, error)
}

// Struct LogReporter is used as ILogReporter interface implementation
type LogReporter struct {
	logReader cw_logReader.ILogReader;
	logParser cw_logParser.ILogParser;
	logAnalyzer cw_logAgnalyzer.ILogAnalyzer;
	errorHandler cw_error.IErrorHandler;
}

type GameReport struct {
	Data orderedMap.OrderedMap
}

type MODGameReport struct {
	Data orderedMap.OrderedMap
}

type IGameReport interface {
	GetData() orderedMap.OrderedMap
}

func (gr GameReport) GetData() orderedMap.OrderedMap {
	return gr.Data;
}

func (mgr MODGameReport) GetData() orderedMap.OrderedMap {
	return mgr.Data
}

//Constructor for LogReporter
func CreateLogReporter(
	logReader cw_logReader.ILogReader,
	logParser cw_logParser.ILogParser,
	logAnalyzer cw_logAgnalyzer.ILogAnalyzer,
	errorHandler cw_error.IErrorHandler) *LogReporter {
		return &LogReporter{
			logReader: logReader,
			logParser: logParser,
			logAnalyzer: logAnalyzer,
			errorHandler: errorHandler,
		}
}

// Func GetGameReport generates the basic game report as requested in CloudWalk Assessment item 3.2
//
// Returns a map where the key is the match number and the value is the match info
func (logReporter *LogReporter) GetGameReport(reportType ReportType) (IGameReport, error) {
	stringLog, readerErr := logReporter.logReader.GetLog()
	parsedLog, parserErr := logReporter.logParser.Parse(stringLog)

	var analyzedMatches IGameReport

	if(readerErr == nil && parserErr == nil) {
		if reportType == GAME_REPORT {
			analyzedMatches = GameReport {
				Data: *logReporter.logAnalyzer.GetMatchesInfo(parsedLog),
			}
		} else {
			analyzedMatches = MODGameReport{
				Data: *logReporter.logAnalyzer.GetMODMatchesInfo(parsedLog),
			}
		}
	}

	var aggregatedErrors = logReporter.errorHandler.BuildErrorOutput(readerErr, parserErr)

	return  analyzedMatches, aggregatedErrors
}