package controllers

import (
	"net/http"
	"text/template"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
)

func LoginWeb(w http.ResponseWriter, r *http.Request) {

	if _, err := auxiliar.ValidaSessao(r); err == nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	

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
