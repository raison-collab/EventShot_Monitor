package main

import (
	hook "github.com/robotn/gohook"
	"time"
)

func StartEventMonitor(config Config, eventChan chan Event) {
	for {
		for _, eventName := range config.Events {
			if hook.AddEvent(eventName) {
				eventChan <- Event{Type: "do_screen", Data: nil}
			}
		}

		time.Sleep(time.Duration(config.Interval) * time.Millisecond) // Задержка для снижения нагрузки на процессор
	}
}
