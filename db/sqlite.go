package db

import (
	"database/sql"
	"log"
)

type SQLiteDB struct {}

func (db SQLiteDB) NewDatabase() *DB {
	dbInstance, err := sql.Open("sqlite3", "./sparrow-plus.db")
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		dbInstance: dbInstance,
	}
}