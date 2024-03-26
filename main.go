package main

import (
	"EventShot_Monitor/config"
	"EventShot_Monitor/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.LoadConfig("config_server.json")
	if err != nil {
		log.Printf("[ERROR] Ошибка обработки конфиг файла: %v\n", err)
		os.Exit(1)
	}

	logFile := ConfigureLogging(cfg.LogFilename)
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

	utils.CreateDir(currentDir + cfg.ScreenshotDir)
	utils.CreateDir(currentDir + cfg.VideoDir)

	// обработчики endpoint
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		UploadHandler(w, r, cfg)
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		PingHandler(w)
	})
	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		VideoHandler(w, r, cfg)
	})
	http.HandleFunc("/register", RegisterHandler)

	log.Printf("[SERVER] start listening")
	log.Fatal(http.ListenAndServe(cfg.ServerUrl, nil))
}
