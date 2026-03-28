package controllers

import (
	"net/http"
	"text/template"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
)

func HomeWeb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		usuario, err := auxiliar.ValidaSessao(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		t, err := template.ParseFiles("templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items := struct {
			NomeUsuario  string
			EmailUsuario string
		}{
			NomeUsuario:  usuario.Nome,
			EmailUsuario: usuario.Email,
		}

		t.Execute(w, items)

	default:
		w.Write([]byte("Erro no metodo"))
	}
}
