package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	// Convert to runes to properly handle Unicode characters
	runes := []rune(s)

	// Replace commas with spaces
	for i, r := range runes {
		if r == ',' {
			runes[i] = ' '
		}
	}

	// Convert back to string and split by whitespace
	return strings.Fields(string(runes))
}
