package controllers

import (
	"net/http"
	"text/template"
)

func HomeWeb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

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

	default:
		w.Write([]byte("Erro no metodo"))
	}
}
