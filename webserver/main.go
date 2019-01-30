package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	http.HandleFunc("/env_all", EnvFullHandler)
	http.HandleFunc("/secret/", SecretHandler)
	http.HandleFunc("/secret_all/", SecretFullHandler)
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/crash", CrashHandler)

	// log ready message
	log.Println("Webserver started")

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
	log.Printf("URL called: %s\n", r.URL.Path[1:])
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

// This is a special case of the Environment Handler, it will iterate
// over all environment variables.
func EnvFullHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL called: %s\n", r.URL.Path[1:])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "List of all environment variables:\n")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Fprintf(w, "%s - %s \n", pair[0], pair[1])
	}
}

// The SecretHandler will try to read a file from the system. Purpose
// is to demonstrate the usage of SecretMaps and the mounting of keys
// as files into the filesystem of a Pod.
func SecretHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL called: %s\n", r.URL.Path[1:])

	path := "/etc/secrets/"
	key := r.URL.Path[8:]
	file := path+key

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(w, "An error occourred reading the specified file: %s\n", file)
			http.Error(w, err.Error(), 500)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Value of file: %s\n", file)
		fmt.Fprintf(w, string(data))
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No secret found for key: %s\n", file)
	}

}

// This is the special case of the above SecretHandler, we will use
// a WalkFunc to list all files in the secret's mount path
func SecretFullHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("URL called: %s\n", r.URL.Path[1:])
	var files []string

	root := "/etc/secrets/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "An error occurred!\n")
		fmt.Fprintf(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	for _, file := range files {
		fmt.Fprintln(w, file)
	}
}