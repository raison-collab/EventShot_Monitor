package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ServerUrl               string   `json:"server_url"`
	ScreenshotPath          string   `json:"screenshot_path"`
	Interval                uint     `json:"interval"`
	Timeout                 uint     `json:"timeout"`
	CompressionLevel        int      `json:"compression_level"`
	MouseClickEvents        []uint16 `json:"mouse_click_events"`
	KeyboardClickEvents     []string `json:"keyboard_click_events"`
	KeyboardClickSpecEvents []uint16 `json:"keyboard_click_spec_events"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config_client.json")
	if err != nil {
		return config, err
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Fatalf("Ошибка закрытия config файла: %v\n", err)
			return
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, err
}
