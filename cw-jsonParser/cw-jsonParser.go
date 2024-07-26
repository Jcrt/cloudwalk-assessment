package cw_jsonParser

import (
	"encoding/json"
	"fmt"
)

func Parse2Json[TValue any](value TValue) string {
	parsedValue, err := json.MarshalIndent(value, "", "  ")
	if(err != nil){
		fmt.Println("Error: ", err)
	}
	return string(parsedValue)
}