package utils

import (
	"EventShot_Monitor/config"
	"log"
	"os"
)

// CreateDir Создаем диррекотрию для хранения скриншотов
func CreateDir(fullPath string) {
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

// GetScreenshotsFilenames Получает имена файлов в дирректории. В дирректории должны быть файлы, лучше всего снаала проверить их наличие
func GetScreenshotsFilenames(config config.Config) ([]string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return []string{}, err
	}

	entries, err := os.ReadDir(currentDir + config.ScreenshotDir)
	if err != nil {
		return []string{}, err
	}
	fileNames := make([]string, 0, len(entries))

	for _, entry := range entries {
		fileNames = append(fileNames, entry.Name())
	}
	return fileNames, nil
}

// HasFilesInScreenshotDir проверяет есть ли файлы в дирректории скриншотов
func HasFilesInScreenshotDir(config config.Config) (bool, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return false, err
	}
	// получаем все файлы из дирректории
	entries, err := os.ReadDir(currentDir + config.ScreenshotDir)
	if err != nil {
		return false, err
	}
	return len(entries) > 0, nil
}
