package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"log"
	"os"
	"time"
)

func StartScreenshotHandler(eventChan chan Event) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка вычисления пути к данной дирректории: %v\n", err)
	}
	createScreensDir(fmt.Sprintf("%s/screens", currentDir))

	for {
		event := <-eventChan
		switch event.Type {
		case "left_click":
			captureScreen(currentDir)
		}

	}
}

// create dir
func createScreensDir(fullPath string) {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			log.Printf("Ошибка создании дирректории: %v\n", err)
		}
	}
}

func captureScreen(currentDir string) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Printf("Ошибка захвата экрана: %v\n", err)
		return
	}
	fileName := fmt.Sprintf("screenshot_%d.png", time.Now().Unix())

	// todo взять путь из конфига
	fullPath := fmt.Sprintf("%s/screens", currentDir)
	// create file
	file, err := os.Create(fmt.Sprintf("%s/%s", fullPath, fileName))
	if err != nil {
		log.Printf("Ошибка создания файла: %v\n", err)
		return
	}
	// close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Ошибка закрытия файла: %v\n", err)
		}
	}(file)

	// check for error
	if png.Encode(file, img) != nil {
		log.Printf("Ошибка записи файла: %v\n", err)
	}

	// todo отпарвить на сервер
}
