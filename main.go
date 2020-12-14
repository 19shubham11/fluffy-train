package main

import (
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox!"))
}


func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("show a snippet!"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create a snippet!"))
}


func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :2021")
	err := http.ListenAndServe(":2021", mux)
	if err != nil {
		log.Fatal("Could not start server on port 2021", err)
	}
}
