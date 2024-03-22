package main

import "log"

type Event struct {
	Type string
	Data interface{}
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка конфигурирования: %v\n", err)
	}

	eventChan := make(chan Event)
	//Запуск монитора событий
	go StartEventMonitor(config, eventChan)

	// Запуск обработчика скриншотов
	StartScreenshotHandler(config, eventChan)
}
