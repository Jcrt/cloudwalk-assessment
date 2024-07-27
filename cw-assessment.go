package main

import (
	cw_consoleApp "cloudwalk-assessment/cw-consoleApp"
	jsonParser "cloudwalk-assessment/cw-jsonParser"
	cw_reporter "cloudwalk-assessment/cw-reporter"
	"fmt"
)


func main() {
	reportType, err := cw_consoleApp.GetApplicationMode()

	if(err != nil) {
		fmt.Println(err.Error())
	} else {
		if(reportType == cw_consoleApp.MODE_GAME_REPORT) {
			gameReport := cw_reporter.GetGameReport()
			fmt.Printf("%s", jsonParser.Parse2Json(gameReport))
		} else if(reportType == cw_consoleApp.MODE_MOD_REPORT) {
			gameReport := cw_reporter.GetMODGameReport()
			fmt.Printf("%s", jsonParser.Parse2Json(gameReport))
		}
	}
}
