package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	"github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func CadastroWeb(w http.ResponseWriter, r *http.Request) {
	
	var store = sessions.NewCookieStore([]byte("chave-secreta"))
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
			session, _ := store.Get(r, "session")
			session.AddFlash("As senhas não coincidem", "erro" )
			session.Save(r,w)
			http.Redirect(w, r, "/cadastro", http.StatusSeeOther)
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

		http.Redirect(w, r, "/login?sucesso=1", http.StatusSeeOther)

	case "GET":
		
		t, err := template.ParseFiles("templates/cadastro.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := store.Get(r, "session")
		erros := session.Flashes("erro")
		sucessos := session.Flashes("sucesso")
		session.Save(r,w)

		items := struct {
		Erros []interface{}
		Sucessos []interface{} 
		}{
			Erros: erros,
			Sucessos: sucessos,
		}

		t.Execute(w, items)
	default:
		w.Write([]byte("Erro no metodo"))
	}
}
