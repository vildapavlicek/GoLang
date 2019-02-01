package handlers

import (
	"html/template"
	"net/http"
)

func SetHandlers(m *http.ServeMux) {
	m.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("./views/index.gohtml"))
	tpl.Execute(w, "Vilda")
}
