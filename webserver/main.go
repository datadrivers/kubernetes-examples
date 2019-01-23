package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	http.HandleFunc("/env/", EnvHandler)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/crash", CrashHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// This is the main handle func of this webserver example, it answers
// to all requests by responding with the text found in the URL path
// after the first forward slash (http://localhost:8080/).
// The example was taken and extended from https://golang.org/doc/articles/wiki/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	log.Printf("URL called: %s\n", r.URL.Path[1:])
}

// Purpose of this HomeHandler is to crash the program by calling it.
// This can be used to show the behaviour when running inside a
// Kubernetes Pod.
func CrashHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	log.Fatal("Crash handler has been called!")
}

// This HomeHandler will try to extract environment keys from the URL
// path and will do an lookup for these keys and print the values.
func EnvHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL called: %s\n", r.URL.Path[1:])
	key := r.URL.Path[5:]
	env, found := os.LookupEnv(key)

	if found {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Key: %s\n", key)
		fmt.Fprintf(w, "Value: %s\n", env)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No environment key found for %s\n", key)
	}
}
