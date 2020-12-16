package main

import (
	"fmt"
	"html/template"
	"log"
    "net/http"
    "strconv"   
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
	}
    
    templateFiles := []string {
        "../../ui/html/home.page.tmpl",
        "../../ui/html/base.layout.tmpl",
        "../../ui/html/footer.partial.tmpl",
    }

	ts, parseErr := template.ParseFiles(templateFiles...)
	if parseErr != nil {
		log.Println("template parse error", parseErr.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	templateErr := ts.Execute(w, nil)
	if templateErr != nil {
		log.Println("template err", parseErr.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("Create a new snippet..."))
}
