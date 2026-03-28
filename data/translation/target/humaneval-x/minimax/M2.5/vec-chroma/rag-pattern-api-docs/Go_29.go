package main

import "strings"

func FilterByPrefix(strings []string, prefix string) []string {
	result := make([]string, 0, len(strings))
	for _, s := range strings {
		if strings.HasPrefix(s, prefix) {
			result = append(result, s)
		}
	}
	return result
}