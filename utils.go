package main

import (
	"log"
	"os"
)

func CreateScreensDir(fullPath string) {
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		err = os.Mkdir(fullPath, os.ModePerm)
		if err != nil {
			log.Printf("Ошибка создании дирректории: %v\n", err)
		}
	}
}
