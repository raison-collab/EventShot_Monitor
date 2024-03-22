package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("screenshot")
	if err != nil {
		http.Error(w, "Error when receiving file", http.StatusBadRequest)
		return
	}

	defer func(file2 multipart.File) {
		err = file2.Close()
		if err != nil {
			http.Error(w, "Error when saving file", http.StatusInternalServerError)
			return
		}
	}(file)

	filePath := filepath.Join("screens", header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error with creating file", http.StatusInternalServerError)
		return
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			http.Error(w, "Error when closing file", http.StatusInternalServerError)
			return
		}
	}(out)

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error with saving file", http.StatusInternalServerError)
		return
	}

	// обработчик скрина
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("File uploaded"))
	if err != nil {
		http.Error(w, "Error when uploading file", http.StatusInternalServerError)
		return
	}
}
