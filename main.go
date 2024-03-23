package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Printf("[ERROR] Ошибка обработки конфиг файла: %v\n", err)
		os.Exit(1)
	}

	logFile := ConfigureLogging(config.LogFilename)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(logFile)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка вычисления пути к данной дирректории: %v\n", err)
	}

	CreateScreensDir(currentDir + config.ScreenshotDir)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		UploadHandler(w, r, config)
	})

	log.Printf("[SERVER] start listening")
	log.Fatal(http.ListenAndServe(config.ServerUrl, nil))
}
