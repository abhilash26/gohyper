package storage

import (
	"github.com/abhilash26/gohyper/internal/option"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitSQLite() error {
	var err error
	databaseURL := option.GetStringEnv("DATABASE_URL", "../database.db")
	DB, err = sqlx.Connect("sqlite3", databaseURL)
	if err != nil {
		return err
	}
	return nil
}
