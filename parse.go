package myeasysql

import (
	"regexp"
	"strings"
)

// _ParseQuerys parses the given SQL SELECT query and extracts the column names.
// It returns a slice of strings containing the column names.
// If the query contains a wildcard (*) for selecting all columns, it returns a slice with a single element "*".
//
// Parameters:
//   - query: A string containing the SQL SELECT query.
//
// Returns:
//   - []string: A slice of strings containing the column names, or nil if the query is invalid or does not match the expected pattern.
func parseQueries(query string) []string {
	step1 := regexp.MustCompile(`SELECT\s+(.*?)\s+FROM`).FindStringSubmatch(query)
	if len(step1) == 0 {
		return nil
	}
	splitted := strings.Split(step1[1], ",")
	for i, s := range splitted {
		if s == "*" {
			return []string{"*"}
		}
		step2 := regexp.MustCompile(".* AS (.*)$").FindStringSubmatch(s)
		if len(step2) != 0 {
			splitted[i] = step2[1]
		} else {
			splitted[i] = strings.TrimSpace(s)
		}
	}
	return splitted
}
