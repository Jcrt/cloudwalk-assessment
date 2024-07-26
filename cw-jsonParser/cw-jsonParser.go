// Package cw_jsonParser is responsible by parse all structs inside the program to Json
package cw_jsonParser

import (
	"encoding/json"
	"fmt"
)

// Func Parse2Json converts structs to Json
// The type constraint TValue is declared to any, which means that any type can be provided
// The parameter value expect any value to be converted to Json
//
// Returns a string containing the json value
func Parse2Json[TValue any](value TValue) string {
	parsedValue, err := json.MarshalIndent(value, "", "  ")
	if(err != nil){
		fmt.Println("Error: ", err)
	}
	return string(parsedValue)
}