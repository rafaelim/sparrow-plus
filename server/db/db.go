package db

import (
	"database/sql"
	"log"
)

type SQLiteDB struct {
	Url string
}

func NewSQLiteStorage(config SQLiteDB) *sql.DB {
	db, err := sql.Open("sqlite3", config.Url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
