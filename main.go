package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	logFile := ConfigureLogging("app.log")
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
	CreateScreensDir(fmt.Sprintf("%s", currentDir+"/screens"))

	http.HandleFunc("/upload", UploadHandler)

	log.Printf("[SERVER] start listening")
	log.Fatal(http.ListenAndServe("127.0.0.1:8082", nil))
}
