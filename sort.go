package myeasysql

import (
	"strings"
)

type ptrToInt struct {
	ptr any
	i   int
}

func _sortPtrs(a, b ptrToInt) bool {
	return a.i < b.i
}

func _SortKeys(tags []string, names []string, keys []string) []int {
	var order []int
	for i := range names {
		for j := range keys {
			var cmpTag = len(tags) != 0 && strings.Compare(tags[i], keys[j]) == 0
			var cmpName = strings.Compare(strings.ToLower(names[i]), strings.ToLower(keys[j])) == 0

			if cmpTag || cmpName {
				order = append(order, j)
			}
		}
	}
	return order
}
