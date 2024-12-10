package myeasysql

import (
	"regexp"
	"strings"
)

func _ParseQuerys(query string) []string {
	step1 := regexp.MustCompile(`SELECT\s+(.*?)\s+FROM`).FindStringSubmatch(query)
	if len(step1) == 0 {
		return nil
	}
	splited := strings.Split(step1[1], ",")
	for i, s := range splited {
		splited[i] = strings.TrimSpace(s)
	}
	return splited
}
