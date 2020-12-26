package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
	"os"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippetForm).Methods("GET")
	r.HandleFunc("/snippet/create", app.createSnippet).Methods("POST")
	r.HandleFunc("/snippet/{id}", app.showSnippet).Methods("GET")

	pwd, err := os.Getwd()
	if err != nil {
		app.errorLog.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir(pwd + "/ui/static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standardMiddleware.Then(r)
}
