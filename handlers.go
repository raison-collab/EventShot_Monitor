package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// todo добавить в конфиг
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка в получении данной директории: %v\n", err)
		return
	}

	// Открываем файл для записи
	file, err := os.Create(fmt.Sprintf("%s/screens/%s", currentDir, r.Header.Get("Filename")))
	if err != nil {
		log.Printf("Ошибка при создании файла: %v\n", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии файла: %v\n", err)
		}
	}(file)

	// Копируем содержимое тела запроса в файл
	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Printf("Ошибка при записи файла: %v\n", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	// обработчик скрина
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("File uploaded"))
	if err != nil {
		log.Printf("ошибка при записи ответа")
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}
