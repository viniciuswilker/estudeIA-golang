package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	"github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func CadastroWeb(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Println("POST NA ROTA DE CADASTRO")
		r.ParseForm()

		nome := r.FormValue("nome")
		sobrenome := r.FormValue("sobrenome")
		email := r.FormValue("email")
		username := r.FormValue("username")
		senha := r.FormValue("senha")
		confirmar_senha := r.FormValue("confirma_senha")

		if senha != confirmar_senha {
			w.Write([]byte("as senhas não coincidem"))
			return
		}

		senhaComHash, err := auxiliar.Hash(senha)
		if err != nil {
			http.Error(w, "Erro ao processar a senha", http.StatusInternalServerError)
			return
		}

		db, err := database.ConectaBanco()
		if err != nil {
			http.Error(w, "Erro ao conectar no banco", 500)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDeUsuarios(db)

		usuario := models.Usuario{
			Username:    username,
			Nome:        nome,
			Sobrenome:   sobrenome,
			Email:       email,
			Senha:       string(senhaComHash),
			TipoUsuario: "A",
		}

		_, err = repositorio.Criar(usuario)
		if err != nil {
			http.Error(w, "Erro ao criar o usuarios", 500)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	case "GET":
		t, err := template.ParseFiles("templates/cadastro.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	default:
		w.Write([]byte("Erro no metodo"))
	}
}
