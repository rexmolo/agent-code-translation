package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}
	
	// Replace commas with spaces, then split by whitespace
	s = strings.ReplaceAll(s, ",", " ")
	return strings.Fields(s)
}
