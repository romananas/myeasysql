package myeasysql

import "database/sql"

type DB struct {
	db *sql.DB
}

// Convert *sql.DB type into myeasysql.DB type
//
// Example Usage:
//
//	func main() {
//		sql_db, _ := sql.Open("sqlite3", "data.do")
//		myeasysql_db := myeasysql.Convert(sql_db)
//		...
//	}
func Convert(db *sql.DB) DB {
	return DB{db: db}
}
