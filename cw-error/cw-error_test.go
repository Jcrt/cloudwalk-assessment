package cw_error

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorHandler = &ErrorHandler{}

func Test_BuildErrorOutput_ShouldReturnNilWhenNoErrors(t *testing.T) {
	errOutput := errorHandler.BuildErrorOutput()

	assert.Nil(t, errOutput)
}

func Test_BuildErrorOutput_ShouldReturnNilWhenErrorsPassedWasNil(t *testing.T) {
	var err1, err2 error

	errOutput := errorHandler.BuildErrorOutput(err1, err2)
	
	assert.Nil(t, errOutput)
}

func Test_BuildErrorOutput_ShouldContainErrorMsgsInsideErrOutput(t *testing.T) {
	err1 := fmt.Errorf("error 1")
	err2 := fmt.Errorf("error 2")

	errOutput := errorHandler.BuildErrorOutput(err1, err2)
	
	assert.NotNil(t, errOutput)

	msg := errOutput.Error()

	assert.Contains(t, msg, "the following error(s) was found:")
	assert.Contains(t, msg, "error 1")
	assert.Contains(t, msg, "error 2")
}