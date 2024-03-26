package main

import (
	config2 "EventShot_Monitor/config"
	"EventShot_Monitor/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	config, err := config2.LoadConfig("config_server.json")
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

	utils.CreateDir(currentDir + config.ScreenshotDir)
	utils.CreateDir(currentDir + config.VideoDir)

	// обработчики endpoint
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		UploadHandler(w, r, config)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		PingHandler(w)
	})
	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		VideoHandler(w, r, config)
	})

	log.Printf("[SERVER] start listening")
	log.Fatal(http.ListenAndServe(config.ServerUrl, nil))
}
