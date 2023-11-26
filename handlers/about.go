package handlers

import (
	"forum/api"
	"html/template"
	"net/http"
)

func AboutHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		api.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		api.BadRequest(w, r)
		return
	}
	tmpl := template.Must(template.ParseFiles("./ui/html/about.html"))
	err := tmpl.ExecuteTemplate(w, "about", nil)
	if err != nil {
		api.InternalError(w, r)
		return
	}
}
