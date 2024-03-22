package main

import (
	"log"
	"net/http"
)

func main() {
	ConfigureLogging("app.log")
	http.HandleFunc("/upload", UploadHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8082", nil))
}
