package main

import (
	"fmt"
	"log"
	"os"
)

type Event struct {
	Type string
	Data interface{}
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка конфигурирования: %v\n", err)
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("[CREATE DIR ERROR] ошибка создания дирректории %s: %v\n\n", config.ScreenshotPath, err)
	}
	CreateScreensDir(fmt.Sprintf("%s", currentDir+config.ScreenshotPath))

	go ScreenshotsSender(config)

	eventChan := make(chan Event)
	//Запуск монитора событий
	go StartEventMonitor(config, eventChan)

	// Запуск обработчика скриншотов
	StartScreenshotHandler(config, eventChan)
}
