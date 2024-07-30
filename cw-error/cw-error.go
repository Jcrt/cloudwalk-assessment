package cw_error

import (
	"fmt"
	"strconv"
)

//Interface IErrorHandler defines all methods offered by error handler
type IErrorHandler interface {
	BuildErrorOutput(errors ...error) error;
}

// Struct ErrorHandler is used as IErrorHandler interface implementation
type ErrorHandler struct {

}

//Func BuildErrorOutput receives an array of errors and concatenate it to show to the
// end user an unique and formatted error
//
// Returns error interface with all errors
func (errorHandler *ErrorHandler) BuildErrorOutput(errors ...error) error{
	var concatedErr error
	var errorString string = ""
	if(len(errors) > 0) {
		for index, item := range errors {
			if(item != nil){
				errorString = errorString + "\n" + strconv.Itoa(index) + ") " + item.Error()
			}
		}

		if(len(errorString) > 0) {
			concatedErr = fmt.Errorf("the following error(s) was found: %s", errorString)
		}
	}

	return concatedErr
}