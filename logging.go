package main

import (
	"log"
	"os"
)

func ConfigureLogging(fileName string) *os.File {
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// setup logs
	log.SetOutput(logFile)

	// add flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Llongfile)

	return logFile
}
