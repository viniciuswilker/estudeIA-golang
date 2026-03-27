package controllers

import (
	"net/http"
	"text/template"
)

func LoginWeb(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func HomeWeb(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items := struct {
		NomeUsuario  string
		EmailUsuario string
	}{
		NomeUsuario:  "vinicius",
		EmailUsuario: "vincius@email.com",
	}

	t.Execute(w, items)

}
