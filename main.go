package main

import (
	"log"
	"net/http"
    "path/filepath"
    "os"
    "io/ioutil"
    "fmt"
    "crypto/rand"
    "mime"
)

const maxUploadSize = 2 * 1024 * 1024 // 2MB
const uploadPath = "./tmp"

func main() {
    http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Println("Server started on localhost:5000, use /upload for uploading files and /files/{fileName} for downloading files.")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func uploadFileHandler() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
        if err := r.ParseMultipartForm(maxUploadSize); err != nil {
            renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
            return
        }

        fileType := r.PostFormValue("type")
        file, _, err := r.FormFile("uploadFile")
        if err != nil {
            renderError(w, "INVALID_FILE", http.StatusBadRequest)
            return
        }
        defer file.Close()

        fileBytes, err := ioutil.ReadAll(file)
        if err != nil {
            renderError(w, "INVALID_FILE", http.StatusBadRequest)
            return
        }

        filetype := http.DetectContentType(fileBytes)
        switch filetype {
        case "text/csv":
        case "image/jpeg", "image/jpg":
        default:
            renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
            return
        }
        fileName := randToken(12)
        fileEndings, err := mime.ExtensionsByType(fileType)
        if err != nil {
            renderError(w, "CAN_NOT_READ_FILE_TYPE", http.StatusInternalServerError)
            return
        }
        newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
        fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

        // write file
        newFile, err := os.Create(newPath)
        if err != nil {
            renderError(w, "CAN_NOT_WRITE_FILE", http.StatusInternalServerError)
            return
        }
        defer newFile.Close()
        if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
            renderError(w, "CAN_NOT_WRITE_FILE", http.StatusInternalServerError)
            return
        }
        w.Write([]byte("SUCCESS"))

    })
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    w.Write([]byte(message))
}

func randToken(len int) string {
    b := make([]byte, len)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
