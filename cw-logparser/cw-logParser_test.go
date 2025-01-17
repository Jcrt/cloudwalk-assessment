package cw_logParser

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
)

type errorHandlerMock struct {
	mock.Mock
}

func (ehm *errorHandlerMock) BuildErrorOutput(errors ...error) error {
	return nil
}

var errHandlerMock = new(errorHandlerMock)

//Tests if log will be processed due the content is empty or not
func Test_Parse_ShouldHaveErrorWhenLogStringIsEmpty(t *testing.T) {
	testcases := []struct {
		input string;
		expected bool;
	} {
		{ "", true },
		{ "    ", true },
		{ "some text", false },
	}

	mockError := fmt.Errorf("log content is empty")
	errHandlerMock.On("BuildErrorOutput", mock.Anything).Return(mockError)
	
	logParser := CreateLogParser(errHandlerMock)

	for index, tc := range testcases {
		log, err := logParser.Parse(tc.input)

		shouldBeNull := (err == nil) == !tc.expected
		shouldContainString := (err != nil) && strings.Contains(err.Error(), "log content is empty")

		if(shouldBeNull && shouldContainString != tc.expected) {
			t.Errorf("%d) Expected to contain: %q, but not found", index, "log content is empty")
		}

		if(len(log.Matches) > 0){
			t.Errorf("Matches length expected: %q. Matches lenght found: %q", "0", strconv.Itoa(len(log.Matches)))
		}
	}
}

//Tests if a match is correctly being created once keyword exists in log regardless string case in log
func Test_Parse_ShouldBeCreated(t *testing.T) {
	testCases := []struct {
		input string;
		matchCounter int;
	} {
		{ "00:00 InitGame: ", 1 }, 
		{ "00:00 INITGAME:", 1 },
		{ "0:00 initgame: ", 1 },
	}

	errHandlerMock.On("BuildErrorOutput", mock.Anything).Return(nil)
	
	logParser := CreateLogParser(errHandlerMock)

	for index, tc := range testCases {
		log, _ := logParser.Parse(tc.input)
		matchesCounter := len(log.Matches)
		if(matchesCounter == 0){
			t.Errorf("%s) Match lenght expected: %s; found: %s; input: %q", 
				strconv.Itoa(index), 
				strconv.Itoa(tc.matchCounter), 
				strconv.Itoa(matchesCounter), 
			  tc.input)
		}
	}
}

//Tests if a kill entry will be processed depending of some situations, as if there is a match or not
func Test_Parse_ShouldCreateAKillEntry(t *testing.T){
	testCases := []struct {
		input string;
		killsCounter int;
	} {
		{ "00:00 InitGame: \n 00:01 Kill: ", 0 }, 
		{ "00:00 InitGame: \n00:01 Kill: 1 2 3:", 1 }, 
		{ "00:00 INITGAME:\n KILL", 0 },
		{ "0:00 kill: ", 0 },
	}

	var killsCounter int

	errHandlerMock.On("BuildErrorOutput", mock.Anything).Return(nil)
	
	logParser := CreateLogParser(errHandlerMock)

	for index, tc := range testCases {
		log, _ := logParser.Parse(tc.input)

		if(len(log.Matches) == 0) {
			killsCounter = 0
		} else {
			killsCounter = len(log.Matches[0].Kills)
		}

		if(killsCounter != tc.killsCounter){
			t.Errorf("%s) Kill lenght expected: %s; found: %s; input: %q", 
				strconv.Itoa(index), 
				strconv.Itoa(tc.killsCounter), 
				strconv.Itoa(killsCounter), 
			  tc.input)
		}
	}
}

func Test_Parse_ShouldParsePlayerCorrectly(t *testing.T){
	testCases := []struct {
		input string;
		PlayersCounter int;
	} {
		{ "00:00 InitGame: \n00:01 clientuserinfochanged: 2 \\nTest ", 0 }, 
		{ "00:00 InitGame: \n00:01 clientuserinfochanged: 2 n\\Test ", 1 }, 
		{ "00:00 InitGame: \n00:01 Kill: 1 2 3:", 0 }, 
		{ "00:00 INITGAME:\n KILL", 0 },
		{ "0:00 kill: ", 0 },
	}

	var playerCounter int

	errHandlerMock.On("BuildErrorOutput", mock.Anything).Return(nil)
	
	logParser := CreateLogParser(errHandlerMock)

	for index, tc := range testCases {
		log, _ := logParser.Parse(tc.input)

		if(len(log.Matches) == 0) {
			playerCounter = 0
		} else {
			playerCounter = len(log.Matches[0].Players)
		}

		if(playerCounter != tc.PlayersCounter){
			t.Errorf("%s) Players lenght expected: %s; found: %s; input: %q", 
				strconv.Itoa(index), 
				strconv.Itoa(tc.PlayersCounter), 
				strconv.Itoa(playerCounter), 
			  tc.input)
		}
	}
}

