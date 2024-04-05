package cmd

import (
	"database/sql"
	"sparrow-plus/db"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func SetupDatabase() *sql.DB {
	dbConfig := db.SQLiteDB{
		Url: "./sparrow-plus.db",
	}

	db := db.NewSQLiteStorage(dbConfig)

	return db
}
