package myeasysql

import "reflect"

func _ReadTags(v any) []string {
	var tags []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		tags = append(tags, field.Tag.Get("sql"))
	}
	return tags
}
