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
	indexTmpl *template.Template
}

func main() {
	app := &App{
		indexTmpl: template.Must(template.ParseFS(files, "templates/layout.html", "templates/index.html")),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.Index)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
