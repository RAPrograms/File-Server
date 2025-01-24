package main

import (
	"embed"
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var templateFiles embed.FS

var templates *template.Template

func main() {
	http.HandleFunc("/", handler)

	files, err := template.New("").ParseFS(templateFiles, "templates/*.html")
	if err != nil {
		log.Panicf("Unable to parse template files because %v", err)
	}

	templates = files

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)

	vd := ViewData{
		Name: "John Smith",
		Widgets: []Widget{
			{"Blue Widget", 12},
			{"Red Widget", 12},
			{"Green Widget", 12},
		}}
	templates.ExecuteTemplate(w, "index.html", vd)
}
