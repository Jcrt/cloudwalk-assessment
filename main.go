package main

import (
	cw_logreader "cloudwalk-assessment/cw-logreader"
	quake3 "cloudwalk-assessment/quake3"
	"fmt"
)

const filepath string = "assets\\qgames.log.txt";

func main() {

	fileContent := cw_logreader.ReadLogFile(filepath)
	fmt.Println(quake3.MOD_BFG)
	fmt.Println(fileContent)
}