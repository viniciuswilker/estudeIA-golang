package repositorios

import (
	"database/sql"

	"github.com/viniciuswilker/estudeIA-golang/internal/models"
)

type fazenda struct {
	db *sql.DB
}

func NovoRepositorioDeFazendas(db *sql.DB) *fazenda {
	return &fazenda{db}
}

func (repositorio fazenda) Cadastro(f models.Fazenda, u models.Usuario) error {
	tx, err := repositorio.db.Begin()
	if err != nil {
		return err
	}

	queryFazenda := "INSERT INTO fazendas (nome, endereco, codigo_fazenda) VALUES (?, ?, ?)"
	resFazenda, err := tx.Exec(queryFazenda, f.Nome, f.Endereco, f.CodigoFazenda)
	if err != nil {
		tx.Rollback()
		return err
	}

	fazendaID, _ := resFazenda.LastInsertId()

	queryUsuario := `INSERT INTO usuarios (username, nome, sobrenome, email, senha, tipo_usuario, fazenda_id) 
					VALUES (?, ?, ?, ?, ?, 'A', ?)`
	resUsuario, err := tx.Exec(queryUsuario, u.Username, u.Nome, u.Sobrenome, u.Email, u.Senha, fazendaID)
	if err != nil {
		tx.Rollback()
		return err
	}

	usuarioID, _ := resUsuario.LastInsertId()

	queryUpdateDono := "UPDATE fazendas SET dono_id = ? WHERE id = ?"
	if _, err := tx.Exec(queryUpdateDono, usuarioID, fazendaID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
