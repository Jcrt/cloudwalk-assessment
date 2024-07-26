package cw_logReader

import (
	"log"
	"os"
)

func ReadLogFile(filepath string) string {
	body, err:= os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(body)
}