package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed templates/*.html templates/layouts/*.html
var files embed.FS

type App struct {
	indexTmpl *template.Template
	db        *sqlx.DB
}

func main() {
	app := &App{
		indexTmpl: template.Must(template.ParseFS(files, "templates/layouts/main.html", "templates/index.html")),
		db:        sqlx.MustConnect("sqlite3", ":memory:"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Index)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
