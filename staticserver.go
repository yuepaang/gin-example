package main

import (
    "fmt"
    "html/template"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "time"
    "log"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type Myhandler struct{}

type home struct {
    Title string
}

const (
    TemplateDir = "./view/"
    UploadDir = "./upload/"
)

func main() {
    server := http.Server{
        Addr: ":5000",
        Handler: &Myhandler{},
        ReadTimeout: 10 * time.Second,
    }
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = index
    mux["/upload"] = upload
    mux["/file"] = StaticServer
    log.Fatal(server.ListenAndServe())
}

func (*Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if hf, ok := mux[r.URL.String()]; ok {
        hf(w, r)
        return
    }
    if ok, _ := regexp.MatchString("/css/", r.URL.String()); ok {
        http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))).ServeHTTP(w, r)
    } else {
        http.StripPrefix("/", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles(TemplateDir + "file.html")
        t.Execute(w, "上传文件")
    } else {
        r.ParseMultipartForm(32 << 20)
        file, handler, err := r.FormFile("uploadFile")
        if err != nil {
            fmt.Fprintf(w, "%v", "上传错误")
            return
        }
        fileExt := filepath.Ext(handler.Filename)
        if !check(fileExt) {
            fmt.Fprintf(w, "%v", "不允许的上传类型")
            return
        }
        fileName := strconv.FormatInt(time.Now().Unix(), 10) + fileExt
        f, _ := os.OpenFile(UploadDir+fileName, os.O_CREATE|os.O_WRONLY, 0660)
        _, err = io.Copy(f, file)
        if err != nil {
            fmt.Fprintf(w, "%v", "上传失败")
            return
        }
        fileDir, _ := filepath.Abs(UploadDir + fileName)
        fmt.Fprintf(w, "%v", fileName+"上传完成，服务器地址： "+fileDir)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    title := home{Title: "首页"}
    t, _ := template.ParseFiles(TemplateDir + "index.html")
    t.Execute(w, title)
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
    http.StripPrefix("/file", http.FileServer(http.Dir("./upload/"))).ServeHTTP(w, r)
}

func check(name string) bool {
    ext := []string{".xlsx", "csv"}

    for _, v := range ext {
        if v == name {
            return false
        }
    }
    return true
}
