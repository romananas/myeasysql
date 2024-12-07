package myeasysql

import "database/sql"

type DB struct {
	db *sql.DB
}

// Convert *sql.DB type into easy_sql.DB type
func Convert(db *sql.DB) DB {
	return DB{db: db}
}
