package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ServerUrl      string   `json:"server_url"`
	ScreenshotPath string   `json:"screenshot_path"`
	Interval       uint     `json:"interval"`
	Events         []string `json:"events"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open("config.json")
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
