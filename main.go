package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/kevindurb/go-htmx/queries"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed templates/*.html templates/layouts/*.html
var files embed.FS

//go:embed schema.sql
var ddl string

type App struct {
	indexTmpl *template.Template
	queries   *queries.Queries
}

func main() {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(ddl); err != nil {
		log.Fatal(err)
	}

	app := &App{
		indexTmpl: template.Must(template.ParseFS(files, "templates/layouts/main.html", "templates/index.html")),
		queries:   queries.New(db),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Index)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
