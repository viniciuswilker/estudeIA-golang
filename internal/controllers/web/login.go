package controllers

import (
	"net/http"
	"text/template"
)

func LoginWeb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items := struct {
		}{}

		t.Execute(w, items)
	default:
		w.Write([]byte("Metodo não permitido"))
	}
}
