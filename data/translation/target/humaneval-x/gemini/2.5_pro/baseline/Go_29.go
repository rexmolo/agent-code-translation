package main

import (
	"strings"
)

// FilterByPrefix filters an input slice of strings only for ones that start with a given prefix.
func FilterByPrefix(stringsSlice []string, prefix string) []string {
	var result []string
	for _, s := range stringsSlice {
		if strings.HasPrefix(s, prefix) {
			result = append(result, s)
		}
	}
	return result
}
