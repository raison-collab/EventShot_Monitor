package main

type Event struct {
	Type string
	Data interface{}
}

func main() {

	eventChan := make(chan Event)
	//Запуск монитора событий
	go StartEventMonitor(eventChan)

	// Запуск обработчика скриншотов
	StartScreenshotHandler(eventChan)
}
