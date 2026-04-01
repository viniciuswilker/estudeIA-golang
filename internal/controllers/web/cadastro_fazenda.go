package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
	"github.com/viniciuswilker/estudeIA-golang/internal/config"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	"github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func CadastroFazendaWeb(w http.ResponseWriter, r *http.Request) {
	const sessaoNome = "cadastro-fazenda"

	switch r.Method {
	case http.MethodGet:
		session, _ := config.Store.Get(r, sessaoNome)
		mensagens := map[string]interface{}{
			"Erros":    session.Flashes("erro"),
			"Sucessos": session.Flashes("sucesso"),
		}
		session.Save(r, w)

		t, err := template.ParseFiles("templates/cadastro-fazenda.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			return
		}

		t.Execute(w, mensagens)

	case http.MethodPost:
		r.ParseForm()

		nomeUsuario := r.FormValue("nome_usuario")
		sobrenomeUsuario := r.FormValue("sobrenome_usuario")
		emailUsuario := r.FormValue("email_usuario")
		senhaUsuario := r.FormValue("senha_usuario")
		confirmaSenha := r.FormValue("confirma_senha")

		usernameUsuario := r.FormValue("username_usuario") // retirar

		codFazenda := r.FormValue("codigo_fazenda")
		nomeFazenda := r.FormValue("nome_fazenda")
		enderecoFazenda := r.FormValue("endereco_fazenda")

		if senhaUsuario != confirmaSenha {
			auxiliar.MensagemFlash(w, r, sessaoNome, "As senhas não coincidem!", "erro", "/cadastro-fazenda")
			log.Println("[AVISO] Cadastro Fazenda: Senhas não coincidem para o email:", emailUsuario)
			return
		}

		senhaComHash, err := auxiliar.Hash(senhaUsuario)
		if err != nil {
			auxiliar.MensagemFlash(w, r, sessaoNome, "Erro interno ao processar senha", "erro", "/cadastro-fazenda")
			log.Println("[ERRO] Falha no Hash de senha:", err)
			return
		}

		fazenda := models.Fazenda{
			Nome:          nomeFazenda,
			Endereco:      enderecoFazenda,
			CodigoFazenda: codFazenda,
		}

		usuario := models.Usuario{
			Nome:      nomeUsuario,
			Sobrenome: sobrenomeUsuario,
			Email:     emailUsuario,
			Username:  usernameUsuario,
			Senha:     string(senhaComHash),
		}

		db, err := database.ConectaBanco()
		if err != nil {
			http.Error(w, "Erro de conexão com o banco", 500)
			return
		}
		defer db.Close()

		repositorio := repositorios.NovoRepositorioDeFazendas(db)

		if err := repositorio.Cadastro(fazenda, usuario); err != nil {
			log.Printf("[ERRO] Falha no cadastro de Fazenda/Dono: %v", err)
			auxiliar.MensagemFlash(w, r, sessaoNome, "Erro ao salvar dados (verifique se o código ou email já existem)", "erro", "/cadastro-fazenda")
			return
		}

		log.Printf("[SUCESSO] Fazenda '%s' criada com sucesso pelo dono '%s'", nomeFazenda, emailUsuario)
		auxiliar.MensagemFlash(w, r, "sessao-estudeia", "Fazenda e Usuário cadastrados com sucesso!", "sucesso", "/")

	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
