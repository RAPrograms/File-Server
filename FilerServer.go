package main

import (
	"embed"
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var basePath string

//go:embed templates
var templateFiles embed.FS

var templates *template.Template

func main() {
	var err error

	basePath, err = filepath.Abs("./")
	if err != nil {
		log.Panicf("Unable to get base path because of %v", err)
	}

	http.HandleFunc("/", handler)

	files, err := template.New("").ParseFS(templateFiles, "templates/*.html")
	if err != nil {
		log.Panicf("Unable to parse template files because %v", err)
	}

	templates = files

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := path.Join(basePath, r.URL.String())
	file, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}

	if !file.IsDir() {
		handle_file_responce(path, w, r)
		return
	}

	handle_directory_responce(path, w, r)
}

func handle_directory_responce(path string, w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handle_file_responce(path string, w http.ResponseWriter, _ *http.Request) {
	file, _ := os.ReadFile(path)

	file_name := path[strings.LastIndex(path, "/")+1:]
	w.Header().Set("Content-Disposition", "inline; filename="+file_name)

	file_type := http.DetectContentType(file)
	w.Header().Set("Content-Type", file_type)

	w.Write(file)
}
