package main

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.tmpl.Execute(w, nil)
}
