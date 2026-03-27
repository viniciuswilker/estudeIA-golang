package database

import (
	"database/sql"

	"github.com/viniciuswilker/estudeIA-golang/internal/config"
)

func ConectaBanco() (*sql.DB, error) {

	tipoBanco := "mysql"

	db, err := sql.Open(tipoBanco, config.StringBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err

}
