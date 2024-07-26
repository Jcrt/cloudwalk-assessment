// Package cw_logReader is the package responsible for gather all log information and provide it to the application
package cw_logReader

import (
	"log"
	"os"
)

// defines the place where log file is located in local machine
const filepath string = "assets\\qgames.log.txt";

// ReadLogFile gets all log information available in log source
// Returns a string containing all log information
func ReadLog() string {
	body, err:= os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read log file: %v", err)
	}
	return string(body)
}