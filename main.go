package main

import (
	"fmt"
	"log"
	"net/http"
)

const maxUploadSize = 2 * 1024 * 1024
const uploadPath = "./tmp"

func main() error {
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Println("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading files.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
