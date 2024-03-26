package config

import (
	"encoding/json"
	"log"
	"os"
)

type UploadFile struct {
	AllowedFileExtensions []string `json:"allowed_file_extensions"`
	AllowedContentTypes   []string `json:"allowed_content_types"`
}

type VideoCfg struct {
	Fps uint `json:"fps"`
}

type Config struct {
	ScreenshotDir string     `json:"screenshot_dir"`
	VideoDir      string     `json:"video_dir"`
	LogFilename   string     `json:"log_filename"`
	ServerUrl     string     `json:"server_url"`
	UploadFiles   UploadFile `json:"upload_files"`
	Video         VideoCfg   `json:"video"`
}

func LoadConfig(cfgFilePath string) (Config, error) {
	var config Config

	file, err := os.Open(cfgFilePath)
	if err != nil {
		return config, err
	}

	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			log.Printf("[ERROR] Ошибка закрытия файла: %v\n", err)
			os.Exit(1)
		}
	}(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, err
}
