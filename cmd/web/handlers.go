package main

import (
	"errors"
	"fmt"
	"github.com/19shubham11/snippetbox/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	latestSnippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}
	for _, snippet := range latestSnippets {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// templateFiles := []string{
	// 	"./ui/html/home.page.tmpl",
	// 	"./ui/html/base.layout.tmpl",
	// 	"./ui/html/footer.partial.tmpl",
	// }

	// ts, parseErr := template.ParseFiles(templateFiles...)
	// if parseErr != nil {
	// 	app.errorLog.Println("template parse error", parseErr.Error())
	// 	app.serverError(w, parseErr)
	// 	return
	// }

	// templateErr := ts.Execute(w, nil)
	// if templateErr != nil {
	// 	app.errorLog.Println("template err", parseErr.Error())
	// 	app.serverError(w, templateErr)
	// 	return
	// }
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	data := &templateData{snippet}
	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// parse template files
	ts, parseErr := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println("error parsing files", parseErr.Error())
		app.serverError(w, parseErr)
		return
	}

	templateErr := ts.Execute(w, data)
	if templateErr != nil {
		app.errorLog.Println("template error", parseErr.Error())
		app.serverError(w, parseErr)
		return
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

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
