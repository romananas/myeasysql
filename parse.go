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
		if strings.Compare(s, "*") == 0 {
			return []string{"*"}
		}
		step2 := regexp.MustCompile(".* AS (.*)$").FindStringSubmatch(s)
		if len(step2) != 0 {
			splited[i] = step2[1]
		} else {
			splited[i] = strings.TrimSpace(s)
		}
	}
	return splited
}
