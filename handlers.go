package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, config Config) {
	if r.Method != http.MethodPost {
		log.Printf("[%s] Not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	filename := r.Header.Get("Filename")

	if !CheckContentType(contentType, config.UploadFiles.AllowedContentTypes) {
		log.Printf("[SERVER] Content-Type %s not allowed", contentType)
		http.Error(w, "Your Content-Type is not allowed, need image/png", http.StatusBadRequest)
		return
	}

	if !CheckFileExtension(filepath.Ext(filename), config.UploadFiles.AllowedFileExtensions) {
		log.Printf("[SERVER] wrong file extension, need .png")
		http.Error(w, "Wrong file extension, need .png", http.StatusBadRequest)
		return
	}

	log.Printf("[%s] %s %s", r.Method, contentType, filename)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("[ERROR] Ошибка в получении данной директории: %v\n", err)
		return
	}

	// Открываем файл для записи
	file, err := os.Create(fmt.Sprintf("%s/%s", currentDir+config.ScreenshotDir, filename))
	if err != nil {
		log.Printf("[ERROR] Ошибка при создании файла: %v\n", err)
		http.Error(w, "[ERROR] Ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("[ERROR] Ошибка при закрытии файла: %v\n", err)
		}
	}(file)

	// Копируем содержимое тела запроса в файл
	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Printf("[ERROR] Ошибка при записи файла: %v\n", err)
		http.Error(w, "[ERROR] Ошибка сервера", http.StatusInternalServerError)
		return
	}

	log.Printf("[SERVER] %s saved", filename)

	// обработчик скрина
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("File uploaded"))
	if err != nil {
		log.Printf("[ERROR] Ошибка при записи ответа")
		http.Error(w, "[ERROR] Ошибка сервера", http.StatusInternalServerError)
		return
	}
	log.Printf("[SERVER] success request %d", http.StatusOK)
}
