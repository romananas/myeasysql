package myeasysql

import (
	"strings"
)

// func _SortKeys(tags []string, names []string, keys []string) []int {
// 	var order []int
// 	for i := range names {
// 		for j := range keys {
// 			var cmpTag = len(tags) != 0 && strings.Compare(tags[i], keys[j]) == 0
// 			var cmpName = strings.Compare(strings.ToLower(names[i]), strings.ToLower(keys[j])) == 0

// 			if cmpTag || cmpName {
// 				order = append(order, j)
// 			}
// 		}
// 	}
// 	return order
// }

func _SortKeys(tags []string, names []string, keys []string) []int {
	var order []int
	for _, key := range keys {
		for i, name := range names {
			var cmpTag = len(tags) != 0 && strings.Compare(tags[i], key) == 0
			var cmpName = strings.Compare(strings.ToLower(name), strings.ToLower(key)) == 0

			if cmpTag || cmpName {
				order = append(order, i)
				break
			}
		}
	}
	return order
}
