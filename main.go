package main

import (
	jsonParser "cloudwalk-assessment/cw-jsonParser"
	cw_logParser "cloudwalk-assessment/cw-logparser"
	"fmt"
)


func main() {
	games := cw_logParser.ParseLog()
	fmt.Println(jsonParser.Parse2Json(games))
}