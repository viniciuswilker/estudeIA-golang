package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/viniciuswilker/estudeIA-golang/internal/auxiliar"
	"github.com/viniciuswilker/estudeIA-golang/internal/database"
	"github.com/viniciuswilker/estudeIA-golang/internal/models"
	repository "github.com/viniciuswilker/estudeIA-golang/internal/repositorios"
)

func LoginAPI(w http.ResponseWriter, r *http.Request) {
	corpoReq, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usuario models.Usuario

	if err := json.Unmarshal(corpoReq, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConectaBanco()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repository.NovoRepositorioDeUsuarios(db)
	fmt.Println("----- CHAMANDO REPO -----")

	usuarioSalvo, err := repositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		log.Fatal(err)
	}

	if err := auxiliar.VerificarSenha(usuarioSalvo.Senha, usuario.Senha); err != nil {
		log.Fatal(err)
	}

	fmt.Println("----- RESPOSTA LOGIN -----")
	fmt.Println(usuarioSalvo)
}
