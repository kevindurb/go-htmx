package main

import (
	"net/http"
	"strconv"

	"github.com/kevindurb/go-htmx/queries"
)

type PageData struct {
	Todos []queries.Todo
}

func (a *App) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := a.queries.ListTodos(r.Context())
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}

	a.todoListTmpl.ExecuteTemplate(w, "main", PageData{
		Todos: todos,
	})
}

func (a *App) CreateTodo(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	desc := r.FormValue("Description")

	if desc == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	a.queries.InsertTodo(r.Context(), queries.InsertTodoParams{
		Description: desc,
		Done:        false,
	})

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *App) MarkTodoDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	a.queries.MarkTodoDone(r.Context(), id)

	http.Redirect(w, r, "/", http.StatusFound)
}
