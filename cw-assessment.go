package main

import (
	cw_consoleApp "cloudwalk-assessment/cw-consoleApp"
	jsonParser "cloudwalk-assessment/cw-jsonParser"
	cw_reporter "cloudwalk-assessment/cw-reporter"
	"fmt"

	"github.com/iancoleman/orderedmap"
)


func main() {
	reportType, err := cw_consoleApp.GetApplicationMode()

	if(err != nil) {
		fmt.Println(err.Error())
	} else {
		var gameReport orderedmap.OrderedMap

		switch reportMode := reportType; reportMode {
			case cw_consoleApp.MODE_GAME_REPORT: {
				gameReport, err = cw_reporter.GetGameReport()
			}
			case cw_consoleApp.MODE_MOD_REPORT: {
				gameReport, err = cw_reporter.GetMODGameReport()
			} 
		}

		if(err != nil){
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Printf("%s", jsonParser.Parse2Json(gameReport))
		}
	}
}
