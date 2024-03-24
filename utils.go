package main

import (
	"log"
	"os"
)

// CreateScreensDir create dir
func CreateScreensDir(fullPath string) {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			log.Printf("Ошибка создании дирректории: %v\n", err)
		}
	}
	log.Printf("[CREATE DIR] Дирректория %s уже существует", fullPath)
}

// GetScreenshotFilename Получает одно имя файла из дирректории скриншотов. В дирректории должен быть файл, для этого сначала слеудет проверить его наличие
func GetScreenshotFilename(config Config) string {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("[DIR ERROR] Ошибка получения текущей дирректории: %v\n\n", err)
		return ""
	}

	entries, err := os.ReadDir(currentDir + config.ScreenshotPath)
	if err != nil {
		log.Printf("[READ DIR ERROR] Ошибка чтения дирректории: %v\n\n", err)
		return ""
	}
	return entries[0].Name()
}

// HasFilesInScreenshotDir проверяет есть ли файла в дирректории скриншотов
func HasFilesInScreenshotDir(config Config) (bool, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return false, err
	}
	// получаем все файлы из дирректории
	entries, err := os.ReadDir(currentDir + config.ScreenshotPath)
	if err != nil {
		return false, err
	}
	return len(entries) > 0, nil
}

func DeleteScreenshot(config Config, fileName string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("[DIR ERROR] Ошибка получения текущей дирректории: %v\n\n", err)
		return
	}

	err = os.Remove(currentDir + config.ScreenshotPath + "/" + fileName)
	if err != nil {
		log.Printf("[FILE REMOVE ERROR] Ошибка удаления файла: %v\n\n", err)
		return
	}
}
