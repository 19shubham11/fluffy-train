package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/19shubham11/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mysqlUsername := os.Getenv("MYSQLUSER")
	mysqlPassword := os.Getenv("MYSQLPASS")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := fmt.Sprintf("%s:%s@/snippetbox?parseTime=true", mysqlUsername, mysqlPassword)
	db, dbErr := openDB(dsn)
	if dbErr != nil {
		errorLog.Fatalf("failed to get mysql pools %v", dbErr)
	}

	defer db.Close()

	templateCache, templateErr := newTemplateCache("./ui/html")
	if templateErr != nil {
		errorLog.Fatal(templateErr)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
