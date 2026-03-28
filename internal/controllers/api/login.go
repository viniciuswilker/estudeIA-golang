package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	"github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	corpoReq, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(corpoReq, &usuario); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	db, err := database.ConectaBanco()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvo, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		http.Error(w, "E-mail ou senha inválidos", http.StatusUnauthorized)
		return
	}
	fmt.Printf("Rota de API solicitada. Usuario: %v\n", usuario.Email)

	if err := auxiliar.VerificarSenha(usuarioSalvo.Senha, usuario.Senha); err != nil {
		http.Error(w, "E-mail ou senha inválidos", http.StatusUnauthorized)
		return
	}

	token, erro := auxiliar.GerarToken(usuarioSalvo.ID, usuarioSalvo.TipoUsuario)
	if erro != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}
	fmt.Printf("USUARIO: %v , ID: %v. Autenticado com sucesso\n", usuario.Email, usuarioSalvo.ID)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   21600,
		SameSite: http.SameSiteLaxMode,
	})
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"id":    strconv.FormatUint(usuarioSalvo.ID, 10),
	})
}
