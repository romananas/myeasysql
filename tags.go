package myeasysql

import "reflect"

// read all tags in the pointer of structure v and return it as a []string
func _ReadTags(v any) []string {
	var tags []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		tags = append(tags, field.Tag.Get("sql"))
	}
	return tags
}
