package myeasysql

import (
	"database/sql"
	"reflect"
)

// Exec executes a query with the given arguments and returns the result.
// It prepares the query, processes the arguments to handle pointers, and
// then executes the statement.
//
// Parameters:
//   - query: The SQL query to be executed.
//   - args: The arguments for the SQL query. Pointers are dereferenced.
//
// Returns:
//   - sql.Result: The result of the query execution.
//   - error: An error if the query execution fails.
func (d DB) Exec(query string, args ...any) (sql.Result, error) {
	var arrArgs []any
	for _, arg := range args {
		rv := reflect.ValueOf(arg)
		if rv.Kind() == reflect.Ptr {
			tmp, err := getPointers(rv)
			if err != nil {
				return nil, err
			}
			arrArgs = append(arrArgs, tmp)
		} else {
			arrArgs = append(arrArgs, arg)
		}
	}
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(arrArgs...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
