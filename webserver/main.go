package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/datadrivers/kubernetes-examples/webserver/statik"
	"github.com/rakyll/statik/fs"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(statikFS)))

	// main page handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/crash", crashHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// This is the main handle func of this webserver example, it answers
// to all requests by responding with the text found in the URL path
// after the first forward slash (http://localhost:8080/).
// The example was taken and extended from https://golang.org/doc/articles/wiki/
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	log.Printf("URL called: %s\n", r.URL.Path[1:])
}

// Purpose of this handler is to crash the program by calling it.
// This can be used to show the behaviour when running inside a
// Kubernetes Pod.
func crashHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Fatal("Crash handler has been called!")
}
