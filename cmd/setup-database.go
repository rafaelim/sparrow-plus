package cmd

import (
	"sparrow-plus/db"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func SetupDatabase() *db.DB {
	var sqlite db.SQLiteDB

	db := sqlite.NewDatabase()
	// db.SetupCategories()

	return db
}