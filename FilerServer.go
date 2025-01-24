package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v", r.URL)
}
