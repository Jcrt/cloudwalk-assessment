package cw_reporter

import (
	cw_logAnalyzer "cloudwalk-assessment/cw-logAnalyzer"
	cw_logParser "cloudwalk-assessment/cw-logParser"
	"fmt"
	"testing"

	orderedMap "github.com/iancoleman/orderedmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type errorHandlerMock struct {
	mock.Mock
}

func (ehm *errorHandlerMock) BuildErrorOutput(errors ...error) error {
	args := ehm.Called(errors)
	return args.Error(0)
}


type logReaderMock struct {
	mock.Mock
}

func (lrm *logReaderMock) GetLog() (string, error) {
	args := lrm.Called()
	return args.String(0), args.Error(1)
}

type logParserMock struct {
	mock.Mock
}

func (lpm *logParserMock) Parse(logString string) (cw_logParser.Log, error) {
	args := lpm.Called(logString)
	return args.Get(0).(cw_logParser.Log), args.Error(1)
}


type logAnalyzerMock struct {
	mock.Mock
}

func (lam *logAnalyzerMock) GetMatchesInfo(parsedLog cw_logParser.Log) *orderedMap.OrderedMap {
	args := lam.Called(parsedLog)
	
	return args.Get(0).(*orderedMap.OrderedMap)
}

func (lam *logAnalyzerMock) GetMODMatchesInfo(parsedLog cw_logParser.Log) *orderedMap.OrderedMap {
	args := lam.Called(parsedLog)
	
	return args.Get(0).(*orderedMap.OrderedMap)
}

func Test_GetGameReport_ShouldReturnMeaningfulData(t *testing.T) {
	logString := "some string"
	matches := make([]cw_logParser.Match, 0)
	players := make(map[int]cw_logParser.Player, 0)
	kills := make([]cw_logParser.Kill, 0)
	matchesInfo := orderedMap.New()
	matchesInfo.Set("game_1", cw_logAnalyzer.MatchInfo{
		TotalKills: 10,
		Players: []string {
			"Ludmilo", "Calabreso", "Taylor Swifto",
		},
	})	

	matches = append(matches, cw_logParser.Match{
		Order: 1,
		Players: players,
		Kills: kills,
	})

	log := cw_logParser.Log{ Matches: matches }

	var errorHandler = new(errorHandlerMock)
	var logReader = new(logReaderMock)
	var logParser = new(logParserMock)
	var logAnalyzer = new(logAnalyzerMock)
	var logReporter ILogReporter = CreateLogReporter(logReader, logParser, logAnalyzer, errorHandler)

	logReader.On("GetLog").Return(logString, nil)
	logParser.On("Parse", logString).Return(log, nil)
	logAnalyzer.On("GetMatchesInfo", log).Return(matchesInfo)
	errorHandler.On("BuildErrorOutput", mock.Anything).Return(nil)
	result, err := logReporter.GetGameReport(GAME_REPORT)

	_, ok := result.(GameReport)

	assert.NotNil(t ,result)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func Test_GetMODGameReport_ShouldReturnMeaningfulData(t *testing.T) {
	logString := "some string"
	matches := make([]cw_logParser.Match, 0)
	players := make(map[int]cw_logParser.Player, 0)
	kills := make([]cw_logParser.Kill, 0)
	killsByMean := make(map[string]int, 0)
	killsByMean["Peixeira"] = 10
	killsByMean["Havaianas de pau"] = 3
	killsByMean["Espada olímpica"] = 2

	matchesInfo := orderedMap.New()
	matchesInfo.Set("game_1", cw_logAnalyzer.MODMatchInfo{
		KillsByMeans: killsByMean,
	})	

	matches = append(matches, cw_logParser.Match{
		Order: 1,
		Players: players,
		Kills: kills,
	})

	log := cw_logParser.Log{ Matches: matches }

	var errorHandler = new(errorHandlerMock)
	var logReader = new(logReaderMock)
	var logParser = new(logParserMock)
	var logAnalyzer = new(logAnalyzerMock)
	var logReporter ILogReporter = CreateLogReporter(logReader, logParser, logAnalyzer, errorHandler)

	logReader.On("GetLog").Return(logString, nil)
	logParser.On("Parse", logString).Return(log, nil)
	logAnalyzer.On("GetMODMatchesInfo", log).Return(matchesInfo)
	errorHandler.On("BuildErrorOutput", mock.Anything).Return(nil)
	result, err := logReporter.GetGameReport(MOD_REPORT)

	_, ok := result.(MODGameReport)

	assert.NotNil(t ,result)
	assert.Nil(t, err)
	assert.True(t, ok)
}


func Test_GetGameReport_ShouldReturnErrorWhenLogParserFails(t *testing.T) {
	errorMsg := "I'm an error"
	logString := "some string"
	matches := make([]cw_logParser.Match, 0)
	err := fmt.Errorf(errorMsg)
	outerErr := fmt.Errorf("I'm an outer error")

	log := cw_logParser.Log{ Matches: matches }

	var errorHandler = new(errorHandlerMock)
	var logReader = new(logReaderMock)
	var logParser = new(logParserMock)
	var logAnalyzer = new(logAnalyzerMock)
	var logReporter ILogReporter = CreateLogReporter(logReader, logParser, logAnalyzer, errorHandler)

	logReader.On("GetLog").Return(logString, err)
	logParser.On("Parse", logString).Return(log, err)

	errorHandler.On("BuildErrorOutput", mock.Anything).Return(outerErr)

	result, err := logReporter.GetGameReport(GAME_REPORT)

	assert.Nil(t ,result)
	assert.NotNil(t, err)
	logAnalyzer.AssertNumberOfCalls(t, "GetMODMatchesInfo", 0)
}