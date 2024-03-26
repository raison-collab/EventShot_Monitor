package main

import (
	"EventShot_Monitor/config"
	"EventShot_Monitor/utils"
	"EventShot_Monitor/video_maker"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, config config.Config) {
	if r.Method != http.MethodPost {
		log.Printf("[%s] Not allowed", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	contentType := r.Header.Get("Content-Type")
	filename := r.Header.Get("Filename")

	if !utils.CheckContentType(contentType, config.UploadFiles.AllowedContentTypes) {
		log.Printf("[SERVER] Content-Type %s not allowed", contentType)
		http.Error(w, "Your Content-Type is not allowed, need image/png", http.StatusBadRequest)
		return
	}

	if !utils.CheckFileExtension(filepath.Ext(filename), config.UploadFiles.AllowedFileExtensions) {
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

func PingHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	log.Printf("[GET] Ping")
}

func VideoHandler(w http.ResponseWriter, r *http.Request, config config.Config) {
	if r.Method != http.MethodGet {
		log.Printf("[POST] %s not allowed", r.URL)
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	err := video_maker.RenderVideo(config)
	if err != nil {
		log.Printf("[RENDER ERROR] Ошибка создания ролика: %v\n\n", err)
		http.Error(w, "Ошибка сервера. Render video error", http.StatusInternalServerError)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("[DIR PATH ERROR] Ошибка формирования путь к данной дирректории: %v\n\n", err)
		http.Error(w, "Ошибка сервера. Render video error", http.StatusInternalServerError)
		return
	}

	filePath := currentDir + config.VideoDir + "/avi_01.avi"
	fileName := "avi_01.avi"

	w.Header().Set("Content-Type", "video/x-msvideo")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")

	http.ServeFile(w, r, filePath)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {}
