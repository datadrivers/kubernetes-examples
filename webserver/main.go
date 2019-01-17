package main

import (
	"fmt"
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(statikFS)))

	// main page handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/crash", crashHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	log.Printf("URL called: %s\n", r.URL.Path[1:])
}

func crashHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Fatal("Crash handler has been called!")
}
