// Package cw_jsonParser is responsible by parse all structs inside the program to Json
package cw_jsonParser

import (
	"encoding/json"
)

//Interface IJsonParser defines all methods offered by json parser
type IJsonParser interface {
	Parse2Json(value any) (string, error);
}

// Struct JsonParser is used as IJsonParser interface implementation
type JsonParser struct {

}

// Func Parse2Json converts structs to Json
// The type constraint TValue is declared to any, which means that any type can be provided
// The parameter value expect any value to be converted to Json
//
// Returns a string containing the json value
func (jsonParser *JsonParser) Parse2Json(value any) (string, error) {
	parsedValue, err := json.MarshalIndent(value, "", "  ")
	if(err != nil) {
		return "", err
	}
	return string(parsedValue), err
}