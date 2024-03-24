package main

import (
	hook "github.com/robotn/gohook"
)

func StartEventMonitor(config Config, eventChan chan Event) {
	evChan := hook.Start()
	defer hook.End()

	for {
		select {
		case ev := <-evChan:
			if isMonitoredEvent(ev, config) {
				eventChan <- Event{Type: "do_screen", Data: nil}
			}
		}
	}
}

// isMonitoredEvent проверяет, является ли событие одним из отслеживаемых
func isMonitoredEvent(ev hook.Event, config Config) bool {
	for _, eventName := range config.MouseClickEvents {
		if ev.Kind == hook.MouseDown && ev.Button == eventName {
			return true
		}
	}
	for _, eventName := range config.KeyboardClickSpecEvents {
		if ev.Kind == hook.KeyHold && ev.Rawcode == eventName {
			return true
		}
	}
	for _, eventName := range config.KeyboardClickEvents {
		if ev.Kind == hook.KeyDown && hook.RawcodetoKeychar(ev.Rawcode) == eventName {
			return true
		}
	}
	return false
}
