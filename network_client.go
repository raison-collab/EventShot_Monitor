package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ScreenshotsSender(config Config) {
	respCodeChan := make(chan int)
	defer close(respCodeChan)

	for {
		hasFiles, err := HasFilesInScreenshotDir(config)
		if err != nil {
			log.Printf("[READ DIR ERROR] Ошибка чтения дирректории: %v\n\n", err)
			continue
		}
		if hasFiles {
			fileName := GetScreenshotFilename(config)

			go func(fName string) {
				// Отправка скриншота на сервер
				respCode, err := SendScreenshotToServer(config, fName)
				if err != nil {
					log.Printf("[SEND ERROR] ошибка отправки файла %s: %v\n\n", fName, err)
					respCodeChan <- -1
					return
				}
				respCodeChan <- respCode
			}(fileName)

			select {
			case respCode := <-respCodeChan:
				if respCode == http.StatusOK {
					DeleteScreenshot(config, fileName)
				} else {
					log.Printf("[NOT SENT] Файл %s не отправлен", fileName)
				}
			case <-time.After(time.Duration(config.Timeout) * time.Second):
				log.Printf("[TIMEOUT] Таймаут отправки файла %s", fileName)
			}
		}
		time.Sleep(time.Duration(config.Interval/2) * time.Millisecond)
	}
}

func SendScreenshotToServer(config Config, fileName string) (int, error) {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка получения пути: %v\n", err)
		return http.StatusInternalServerError, err
	}

	fullPath := fmt.Sprintf("%s/%s", currentDir+config.ScreenshotPath, fileName)
	file, err := os.Open(fullPath)
	if err != nil {
		log.Printf("Ошибка открытия файла: %v\n", err)
		return http.StatusInternalServerError, err
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
		return http.StatusInternalServerError, err
	}

	// Create request
	req, err := http.NewRequest(http.MethodPost, config.ServerUrl, buffer)
	if err != nil {
		log.Printf("Ошибка формирования запроса: %v\n", err)
		return http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "image/jpeg")
	req.Header.Set("Filename", fileName)

	// Send Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка отправки запроса: %v\n", err)
		return http.StatusInternalServerError, err
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
		return http.StatusInternalServerError, err
	}

	log.Println("Успешно отправлен")
	return http.StatusOK, nil
}
