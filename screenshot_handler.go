package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/jpeg"
	"log"
	"os"
	"time"
)

func StartScreenshotHandler(config Config, eventChan chan Event) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Ошибка вычисления пути к данной дирректории: %v\n", err)
	}

	for {
		select {
		case event := <-eventChan:
			switch event.Type {
			case "do_screen":
				captureScreen(config, currentDir)
			}
		}
	}
}

func captureScreen(config Config, currentDir string) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Printf("Ошибка захвата экрана: %v\n", err)
		return
	}
	fileName := fmt.Sprintf("screen_%d.jpg", time.Now().Unix())

	fullPath := fmt.Sprintf("%s", currentDir+config.ScreenshotPath)
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

	var opts jpeg.Options
	opts.Quality = config.CompressionLevel // 50%
	// check for error
	if jpeg.Encode(file, img, &opts) != nil {
		log.Printf("Ошибка записи файла: %v\n", err)
	}

	//go SendScreenshotToServer(config, fileName)
}
