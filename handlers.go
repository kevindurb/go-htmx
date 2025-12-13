package main

import (
	"net/http"
)

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.indexTmpl.ExecuteTemplate(w, "main", nil)
}
