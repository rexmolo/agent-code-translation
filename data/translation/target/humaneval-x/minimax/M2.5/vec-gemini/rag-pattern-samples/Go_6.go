package main

import (
	"fmt"
	"strings"
)

func ParseNestedParens(paren_string string) []int {
	// Split by spaces and filter empty strings
	parts := strings.Fields(paren_string)

	result := make([]int, 0, len(parts))

	for _, s := range parts {
		depth := 0
		maxDepth := 0
		for _, c := range s {
			if c == '(' {
				depth++
				if depth > maxDepth {
					maxDepth = depth
				}
			} else { // c == ')'
				depth--
			}
		}
		result = append(result, maxDepth)
	}

	return result
}

func main() {
	result := ParseNestedParens("(()()) ((())) () ((())()())")
	fmt.Println(result)
}
