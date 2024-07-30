package cw_consoleApp

import (
	"fmt"
	"os"
)

// ReportType expose all avaliable report types
type applicationMode string
const (
	MODE_GAME_REPORT applicationMode = "game"
	MODE_MOD_REPORT applicationMode = "mod"
	MODE_HELP applicationMode = "help"
)

//Interface IConsoleApp defines all methods offered by user interface
type IConsoleApp interface {
	GetApplicationMode() (applicationMode, error)
}

// Struct ConsoleApp is used as IConsoleApp interface implementation
type ConsoleApp struct {

}

// Func GetApplicationMode shows a menu to the user can input arguments in the program execution
//
// Returns applicationMode that defines which type of report will be generated and error
func (consoleApp ConsoleApp) GetApplicationMode() (applicationMode, error) {
	var mode applicationMode
	var err error = nil

	if(len(os.Args) > 1) {
		firstArg := os.Args[1]

		if(len(firstArg) > 0) {
			switch option := firstArg; option {
				case "help": {
					fmt.Println("CloudWalk Assessment help center ===============")						
					fmt.Println("Available arguments ")
					fmt.Println("    -t <report type> The type of desired report output; OPTIONS: game, mod")
					fmt.Println("================================================")
					mode = MODE_HELP
				}

				case "-t": {
					if(len(os.Args) > 2) {
						secondArg := os.Args[2]

						if(secondArg == string(MODE_GAME_REPORT)){
							mode = MODE_GAME_REPORT
						} else if(secondArg == string(MODE_MOD_REPORT)) {
							mode = MODE_MOD_REPORT
						} else {
							err = fmt.Errorf("invalid type option: %s", secondArg)
						}
					} else {
						err = fmt.Errorf("type wasn't provided after -t argument")
					}
				}
				default: {
					err = fmt.Errorf("invalid argument: %s", firstArg)
				}
			}
		} else {
			mode = MODE_GAME_REPORT
		}
	} else {
		mode = MODE_GAME_REPORT
	}
	
	return mode, err
}