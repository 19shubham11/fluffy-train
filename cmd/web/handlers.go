package main

import (
	"errors"
	"fmt"
	"github.com/19shubham11/snippetbox/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	latestSnippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: latestSnippets,
	})
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: snippet,
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// dummy snippet for now

	title := "Oh look"
	content := "Oh look again, I am trying to insert this, such post much wow"
	expires := "7"

	id, dbErr := app.snippets.Insert(title, content, expires)
	if dbErr != nil {
		app.serverError(w, dbErr)
		return
	}

	res := fmt.Sprintf("Created a new snippet with id %d", id)
	w.Write([]byte(res))
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create snippet"))
}
