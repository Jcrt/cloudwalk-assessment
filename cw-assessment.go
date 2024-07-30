package main

import (
	cw_consoleApp "cloudwalk-assessment/cw-consoleApp"
	cw_error "cloudwalk-assessment/cw-error"
	jsonParser "cloudwalk-assessment/cw-jsonParser"
	cw_logAnalyzer "cloudwalk-assessment/cw-logAnalyzer"
	cw_logParser "cloudwalk-assessment/cw-logParser"
	cw_logReader "cloudwalk-assessment/cw-logReader"
	cw_reporter "cloudwalk-assessment/cw-reporter"
	"fmt"
)


func main() {
	var logReader cw_logReader.ILogReader = &cw_logReader.LogReader{}
	var logParser cw_logParser.ILogParser = &cw_logParser.LogParser{}
	var logAnalyzer cw_logAnalyzer.ILogAnalyzer = &cw_logAnalyzer.LogAnalyzer{}
	var errorHandler cw_error.IErrorHandler = &cw_error.ErrorHandler{}
	var consoleApp cw_consoleApp.IConsoleApp = &cw_consoleApp.ConsoleApp{}

	logReporter := cw_reporter.CreateLogReporter(logReader, logParser, logAnalyzer, errorHandler)

	applicationMode, err := consoleApp.GetApplicationMode()

	if(err != nil) {
		fmt.Println(err.Error())
	} else {
		gameReport, err := logReporter.GetGameReport(cw_reporter.ReportType(applicationMode))

		if(err != nil){
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Printf("%s", jsonParser.Parse2Json(gameReport))
		}
	}
}
