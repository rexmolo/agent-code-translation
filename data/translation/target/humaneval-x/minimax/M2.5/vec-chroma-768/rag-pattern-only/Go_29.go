package main

import "strings"

// FilterByPrefix filters an input slice of strings for ones that start with a given prefix.
// Returns a new slice containing only strings that have the given prefix.
func FilterByPrefix(strings []string, prefix string) []string {
	var result []string
	for _, s := range strings {
		if strings.HasPrefix(s, prefix) {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	// Example usage for testing
	empty := FilterByPrefix([]string{}, "a")
	_ = empty
	
	filtered := FilterByPrefix([]string{"abc", "bcd", "cde", "array"}, "a")
	_ = filtered
}
