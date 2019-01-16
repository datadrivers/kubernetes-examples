package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

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
