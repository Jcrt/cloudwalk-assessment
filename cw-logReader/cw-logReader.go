// Package cw_logReader is the package responsible for gather all log information and provide it to the application
package cw_logReader

import (
	"fmt"
	"log"
	"os"
)

// defines the place where log file is located in local machine
const filepath string = `assets\qgames.log.txt`;

// stores value of single instance of singleStringLog
var stringLogSingleton *singleStringLog

// singleStringLog is used to store the string log during application execution to save memory consumption
type singleStringLog struct {
	stringLog string
}

//Func GetLog provides the string log of the application
// 
// Returns the string containing log and error
func GetLog() (string, error) {
	logObject, err := getInstance() 
	return logObject.stringLog, err
}

// Func getInstance try to retrieve the instance of singleStringLog. 
// If it's null, retrieves the log from file and store it, if not null, returns the current instance
//
// Returns singleStringLog pointer with log string and error
func getInstance() (*singleStringLog, error) {
	var err error
	if(stringLogSingleton == nil){
		logText, err := readLog()

		if(err == nil) {
			stringLogSingleton = &singleStringLog{
				stringLog: logText,
			}
		}
	}

	return stringLogSingleton, err
}

// Func readLog gets all log information available in log source
// Returns a string containing all log information and error
func readLog() (string, error) {
	body, err:= os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read log file: %v", err)
		err = fmt.Errorf("Problems during log file read operation.")
	}
	return string(body), err
}