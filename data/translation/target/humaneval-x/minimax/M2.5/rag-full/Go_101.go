package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	s = strings.ReplaceAll(s, ",", " ")
	return strings.Fields(s)
}
