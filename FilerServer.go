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
	}

	handle_directory_responce(path, w, r)
}

func handle_directory_responce(path string, w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func handle_file_responce(path string, w http.ResponseWriter, r *http.Request) {
	log.Printf("Responding with %s content", path)
}
