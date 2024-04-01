package db

import "database/sql"

type DB struct {
	dbInstance *sql.DB
}

type DatabaseBuilder interface {
	NewDatabase() DB
}