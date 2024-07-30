package cw_jsonParser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonParser IJsonParser = &JsonParser{

}

func Test_Parse2Json_ShouldParseToJsonWithoutErrors(t *testing.T){
	data := struct {
		Prop1, Prop2 string
	} {
		Prop1: "I am",
		Prop2: "a Json now",
	}

	output, _ := jsonParser.Parse2Json(data)

	assert.JSONEq(t, `{
		"Prop1": "I am",
		"Prop2": "a Json now"
	}`, output)
}
