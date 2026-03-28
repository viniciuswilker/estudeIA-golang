package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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
