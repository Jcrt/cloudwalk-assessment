package main

import (
	jsonParser "cloudwalk-assessment/cw-jsonParser"
	cw_reporter "cloudwalk-assessment/cw-reporter"
	"fmt"
)


func main() {
	// gameReport := cw_reporter.GetGameReport()
	// fmt.Printf(jsonParser.Parse2Json(gameReport))

	MODgameReport := cw_reporter.GetMODGameReport()
	fmt.Printf(jsonParser.Parse2Json(MODgameReport))
}