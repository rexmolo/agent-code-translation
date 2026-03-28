package main

import (
	"fmt"
	"strings"
)

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
	// Test cases
	fmt.Println(FilterByPrefix([]string{}, "a"))
	fmt.Println(FilterByPrefix([]string{"abc", "bcd", "cde", "array"}, "a"))
}
