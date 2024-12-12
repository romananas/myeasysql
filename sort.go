package myeasysql

import (
	"strings"
)

// _SortKeys reorders indices of the `names` array (or `tags`, if available) based on their matching order with `keys`.
// It ensures that the returned indices correspond to the order of elements in `keys`.
//
// Parameters:
// - tags: (optional) A slice of strings representing tags associated with `names`.
// - names: A slice of strings representing the names to be matched.
// - keys: A slice of strings representing the desired order of matching keys.
//
// How it works:
// 1. Iterates over each element in `keys`.
// 2. For each key, iterates over `names` (and `tags` if provided).
// 3. Checks if either a tag (`tags[i]`) or a name (`names[i]`) matches the current key (case-insensitive for names).
// 4. If a match is found, appends the index `i` of the matching name or tag to the result list `order`.
// 5. Breaks out of the inner loop after finding the first match for the current key to ensure correct ordering.
//
// Returns:
// - A slice of integers representing the reordered indices of `names` (or `tags`) based on the order in `keys`.
//
// Example:
//
//	tags := []string{"id", "username", "birth", "password"}
//	names := []string{"id", "username", "birth", "password"}
//	keys := []string{"username", "password", "birth"}
//
//	result := _SortKeys(tags, names, keys)
//	// Result: [1, 3, 2]
//	// Explanation: "username" maps to index 1, "password" to index 3, and "birth" to index 2.
func _SortKeys(tags []string, names []string, keys []string) []int {
	var order []int
	nameIndex := make(map[string]int)
	tagIndex := make(map[string]int)

	for i, name := range names {
		nameIndex[strings.ToLower(name)] = i
		if len(tags) != 0 && i < len(tags) {
			tagIndex[tags[i]] = i
		}
	}

	for _, key := range keys {
		if idx, found := tagIndex[key]; found {
			order = append(order, idx)
		} else if idx, found := nameIndex[strings.ToLower(key)]; found {
			order = append(order, idx)
		}
	}

	return order
}
