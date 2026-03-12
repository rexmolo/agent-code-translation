package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	// Replace commas with spaces, then split on whitespace
	return strings.Fields(strings.ReplaceAll(s, ",", " "))
}