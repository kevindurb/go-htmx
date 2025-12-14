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
	todoListTmpl *template.Template
	queries      *queries.Queries
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
		todoListTmpl: template.Must(template.ParseFS(files, "templates/layouts/main.html", "templates/todo-list.html")),
		queries:      queries.New(db),
	}

	mux := http.NewServeMux()
	mux.Handle("GET /", http.RedirectHandler("/todos", http.StatusFound))
	mux.HandleFunc("GET /todos", app.ListTodos)
	mux.HandleFunc("POST /todos", app.CreateTodo)
	mux.HandleFunc("POST /todos/{id}/done", app.MarkTodoDone)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
