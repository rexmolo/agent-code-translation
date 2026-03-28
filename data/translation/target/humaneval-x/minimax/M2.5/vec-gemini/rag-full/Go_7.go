package main

import (
	"fmt"
	"strings"
)

func FilterBySubstring(stringList []string, substring string) []string {
	result := make([]string, 0, len(stringList))
	for _, s := range stringList {
		if strings.Contains(s, substring) {
			result = append(result, s)
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(FilterBySubstring([]string{}, "a"))
	fmt.Println(FilterBySubstring([]string{"abc", "bacd", "cde", "array"}, "a"))
}
