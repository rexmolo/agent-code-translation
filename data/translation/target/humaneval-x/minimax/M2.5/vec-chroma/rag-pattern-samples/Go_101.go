package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	s = strings.Replace(s, ",", " ", -1)
	return strings.Fields(s)
}