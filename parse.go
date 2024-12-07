package myeasysql

import (
	"regexp"
	"strings"
)

func _ParseQuerys(query string) []string {
	re := regexp.MustCompile(`SELECT\s+(.*?)\s+FROM`)
	match := re.FindStringSubmatch(query)
	if len(match) == 0 {
		return nil
	}
	splited := strings.Split(match[1], ",")
	for i, s := range splited {
		splited[i] = strings.TrimSpace(s)
	}
	return splited
}
