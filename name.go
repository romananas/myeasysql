package myeasysql

import "reflect"

func _ReadNames(v any) []string {
	var names []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		names = append(names, field.Name)
	}
	return names
}
