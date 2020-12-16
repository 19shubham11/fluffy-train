package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet", showSnippet)
    mux.HandleFunc("/snippet/create", createSnippet)


    fileServer := http.FileServer(http.Dir("../../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Starting server on :2022")
    err := http.ListenAndServe(":2022", mux)
    log.Fatal(err)
}