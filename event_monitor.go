package main

import (
	hook "github.com/robotn/gohook"
	"time"
)

func StartEventMonitor(eventChan chan Event) {
	for {
		if hook.AddEvent("mleft") {
			eventChan <- Event{Type: "left_click", Data: nil}
		}

		// todo сделать новую логику задержки, основанную на конфиге
		time.Sleep(500 * time.Millisecond) // Задержка для снижения нагрузки на процессор
	}
}
