package controllers

import (
	"fmt"
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
		fmt.Printf("Entrou na Home. Usuario: %v\n", usuario.Username)

		t, err := template.ParseFiles("templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items := struct {
			NomeUsuario  string
			EmailUsuario string
			TituloPagina string
			Titulo       string
		}{
			NomeUsuario:  usuario.Nome,
			EmailUsuario: usuario.Email,
			TituloPagina: "Dashboard",
			Titulo:       fmt.Sprintf("Bem vindo, %s", usuario.Nome),
		}

		t.Execute(w, items)

	default:
		w.Write([]byte("Erro no metodo"))
	}
}
