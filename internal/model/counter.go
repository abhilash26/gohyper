package model

import (
	"github.com/jmoiron/sqlx"
)

type Counter struct {
	ID    int `db:"id"`
	Value int `db:"value"`
}

func InitCounterTable(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS counter (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			value INTEGER
		)`)
	_, err = db.Exec(`INSERT OR IGNORE INTO counter (value) 
    VALUES (?)`, 0)
	return err
}

func GetCounter(db *sqlx.DB) (int, error) {
	var counter Counter
	err := db.Get(&counter, "SELECT * FROM counter LIMIT 1")
	if err != nil {
		return 0, err
	}
	return counter.Value, nil
}

func UpdateCounterValue(db *sqlx.DB, value int) error {
	_, err := db.Exec("UPDATE counter SET value = ? WHERE id = 1", value)
	return err
}
