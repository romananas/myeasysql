package myeasysql

import "database/sql"

type DB struct {
	db *sql.DB
}

// Convert takes a pointer to an sql.DB object and returns a DB struct
// initialized with the provided sql.DB object.
//
// Parameters:
//   - db: A pointer to an sql.DB object representing the database connection.
//
// Returns:
//   - A DB struct with the db field set to the provided sql.DB object.
func Convert(db *sql.DB) DB {
	return DB{db: db}
}
