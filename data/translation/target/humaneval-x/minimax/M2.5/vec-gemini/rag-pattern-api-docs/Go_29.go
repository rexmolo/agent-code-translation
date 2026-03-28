package main

import (
	"strings"
)

// FilterByPrefix filters an input list of strings only for ones that start with a given prefix.
func FilterByPrefix(strings []string, prefix string) []string {
	result := make([]string, 0, len(strings))
	for _, s := range strings {
		if strings.HasPrefix(s, prefix) {
			result = append(result, s)
		}
	}
	return result
}