package cmd

import (
	"database/sql"
	"sparrow-plus/db"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func SetupDatabase(dbUrl string) *sql.DB {
	dbConfig := db.SQLiteDB{
		Url: dbUrl,
	}

	db := db.NewSQLiteStorage(dbConfig)

	return db
}
