package main

import (
	"log"
	"net/http"
)

type PageData struct {
	todos []Todo
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}

	err := a.db.Select(&todos, "SELECT id, description, done, created FROM todos")
	if err != nil {
		log.Fatal("Could not select todos")
	}

	a.indexTmpl.ExecuteTemplate(w, "main", PageData{
		todos,
	})
}
