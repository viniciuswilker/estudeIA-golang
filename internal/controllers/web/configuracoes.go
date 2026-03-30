package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
)

func ConfiguracoesWeb(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		usuario, err := auxiliar.ValidaSessao(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Printf("Entrou nas Configurações. Usuario: %v\n", usuario.Username)

		t, err := template.ParseFiles("templates/configuracoes.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items := struct {
		}{}

		t.Execute(w, items)

	default:
		w.Write([]byte("metodo não permitido"))

	}

}
