package repository

import (
	"database/sql"

	"github.com/viniciuswilker/estudeIA-golang/internal/models"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {

	smtm, err := repositorio.db.Prepare("insert into usuarios (username, nome, sobrenome, email, senha, tipo_usuario) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, nil
	}
	defer smtm.Close()

	resultado, err := smtm.Exec(usuario.Username, usuario.Nome, usuario.Sobrenome, usuario.Email, usuario.Senha, usuario.TipoUsuario)
	if err != nil {
		return 0, nil
	}

	ultimoId, err := resultado.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return uint64(ultimoId), nil

}
