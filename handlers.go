package main

import (
	"log"
	"net/http"

	"github.com/kevindurb/go-htmx/queries"
)

type PageData struct {
	todos []queries.Todo
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	todos, err := a.queries.ListTodos(r.Context())
	if err != nil {
		log.Fatal("Could not select todos")
	}

	a.indexTmpl.ExecuteTemplate(w, "main", PageData{
		todos,
	})
}
