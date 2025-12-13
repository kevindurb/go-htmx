package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates/*.html
var files embed.FS

type App struct {
	tmpl *template.Template
}

func main() {
	app := &App{
		tmpl: template.Must(template.ParseFS(files, "templates/*.html")),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Index)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.tmpl.Execute(w, nil)
}
