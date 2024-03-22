package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func SendScreenshotToServer(config Config, fileName string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка получения пути: %v\n", err)
		return
	}

	fullPath := fmt.Sprintf("%s/%s", currentDir+config.ScreenshotPath, fileName)
	file, err := os.Open(fullPath)
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
		return
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("Ошибка закрытия файла: %v\n", err)
			return
		}
	}(file)

	// Read file
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		log.Printf("ошибка чтения файла: %v\n", err)
		return
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, config.ServerUrl, buffer)
	if err != nil {
		log.Printf("Ошибка формирования запроса: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "image/png")
	req.Header.Set("Filename", fileName)

	// Send Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка отправки запроса: %v\n", err)
		return
	}
	defer func(r *http.Response) {
		err = r.Body.Close()
		if err != nil {
			log.Printf("Ошибка закрытия Body %v\n", err)
			return
		}
	}(resp)

	// Check status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Ошибка. Сервер ответил с кодом %d", resp.StatusCode)
		return
	}

	log.Println("Успешно отправлен")
}
