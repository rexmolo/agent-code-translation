package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	// Replace commas with spaces, then split by whitespace
	return strings.Fields(strings.ReplaceAll(s, ",", " "))
}
