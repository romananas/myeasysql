package myeasysql

import "reflect"

// _ReadTags extracts and returns a slice of SQL tags from the struct fields of the provided value.
// The function takes an interface{} as input, which should be a pointer to a struct.
// It uses reflection to iterate over the struct fields and collect the values of the "sql" tags.
//
// Parameters:
//   - v: an interface{} that should be a pointer to a struct.
//
// Returns:
//   - []string: a slice containing the values of the "sql" tags from the struct fields.
func _ReadTags(v any) []string {
	var tags []string
	tv := reflect.TypeOf(v).Elem()
	for i := range tv.NumField() {
		field := tv.Field(i)
		tags = append(tags, field.Tag.Get("sql"))
	}
	return tags
}
