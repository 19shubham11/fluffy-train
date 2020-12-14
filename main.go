package main

import (
	"log"
	"net/http"
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Starting server on :2021")
	err := http.ListenAndServe(":2021", mux)
	if err != nil {
		log.Fatal("Could not start server on port 2021", err)
	}
}
