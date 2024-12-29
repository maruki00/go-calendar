package pkg

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

const (
	TRIES   = 3
	TIMEOUT = 30
)

type DBHandler struct {
	db *sql.DB
}

func NewDBHandler(dsn string) *DBHandler {
	counter := 0
	var cnx *sql.DB
	var err error
	for counter < TRIES {
		counter++
		slog.Info("try to connect to database.")
		cnx, err = sql.Open("sqlite3", dsn)
		if err != nil {
			continue
		}
	}

	if cnx == nil {
		panic("Could not connect to database!")
	}

	return &DBHandler{
		db: cnx,
	}

}

func (db *DBHandler) GetDB() *sql.DB {
	return db.db
}
