package repositorios

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

func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {

	row, err := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer row.Close()

	var usuario models.Usuario

	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, err
		}
	}

	return usuario, nil

}

func (repositorios usuarios) BuscarPorID(id uint64) (models.Usuario, error) {

	row, err := repositorios.db.Query("select id, nome, sobrenome, username, email, criadoEm from usuarios where id = ?", id)
	if err != nil {
		return models.Usuario{}, err
	}

	defer row.Close()

	var usuario models.Usuario

	if row.Next() {

		if err := row.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Sobrenome,
			&usuario.Username,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}

	}
	return usuario, nil

}
