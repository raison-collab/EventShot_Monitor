package main

import (
	"log"
	"os"
)

// CreateScreensDir Создаем диррекотрию для хранения скриншотов
func CreateScreensDir(fullPath string) {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			log.Printf("Ошибка создании дирректории: %v\n", err)
		}
		log.Printf("[SERVER] %s дирректория создана", fullPath)
	}
	log.Printf("[SERVER] %s Уже существует", fullPath)
}

// CheckContentType Проверяем допустимый ли заголовок Content-Type
func CheckContentType(current string, allowed []string) bool {
	for _, ct := range allowed {
		if ct == current {
			return true
		}
	}
	return false
}

// CheckFileExtension Проверяем расширение файла
func CheckFileExtension(current string, allowed []string) bool {
	for _, ct := range allowed {
		if ct == current {
			return true
		}
	}
	return false
}
