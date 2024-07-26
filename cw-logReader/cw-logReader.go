// Package cw_logReader is the package responsible for gather all log information and provide it to the application
package cw_logReader

import (
	"log"
	"os"
)

// defines the place where log file is located in local machine
const filepath string = "assets\\qgames.log.txt";

// stores value of single instance of singleStringLog
var stringLogSingleton *singleStringLog

// singleStringLog is used to store the string log during application execution to save memory consumption
type singleStringLog struct {
	stringLog string
}

//Func GetLog provides the string log of the application
func GetLog() string {
	return getInstance().stringLog
}

// Func getInstance try to retrieve the instance of singleStringLog. 
// If it's null, retrieves the log from file and store it, if not null, returns the current instance
func getInstance() *singleStringLog {
	if(stringLogSingleton == nil){
		stringLogSingleton = &singleStringLog{
			stringLog: readLog(),
		}
	}

	return stringLogSingleton
}

// ReadLogFile gets all log information available in log source
// Returns a string containing all log information
func readLog() string {
	body, err:= os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read log file: %v", err)
	}
	return string(body)
}