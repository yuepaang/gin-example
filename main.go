package main

import (
    "fmt"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

const UPLOADFOLDER = "./upload/"

func main() {
    router := gin.Default()
    router.MaxMultipartMemory = 8 << 20 // 8MB
    router.Static("/", "./templates")
    router.POST("/upload", func(c *gin.Context) {
        form, err := c.MultipartForm()
        if err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
            return
        }
        files := form.File["files"]

        for _, file := range files {
            filename := filepath.Base(file.Filename)
            if err := c.SaveUploadedFile(file, UPLOADFOLDER+filename); err != nil {
                c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
                return
            }
        }

        c.Header("Content-Type", "application/json")
        c.JSON(http.StatusOK, gin.H {
            "message": fmt.Sprintf("Uploaded successfully %d files", len(files)),
        })
        // c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files", len(files)))
    })
    router.Run(":5000")
}
