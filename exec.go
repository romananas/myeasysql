package myeasysql

import (
	"database/sql"
	"reflect"
)

func (d DB) Exec(query string, args ...any) (sql.Result, error) {
	var arrArgs []any
	for _, arg := range args {
		var rv = reflect.ValueOf(arg)
		if isPtr(rv) {
			tmp, err := getPointers(arg)
			if err != nil {
				return nil, err
			}
			arrArgs = append(arrArgs, tmp...)
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
