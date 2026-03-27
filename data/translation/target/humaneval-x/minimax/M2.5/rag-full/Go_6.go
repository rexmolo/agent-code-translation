package main

import (
	"fmt"
	"strings"
)

func ParseNestedParens(paren_string string) []int {
	// Split by space and filter empty strings
	groups := strings.Split(paren_string, " ")

	var result []int

	for _, group := range groups {
		if group == "" {
			continue
		}

		depth := 0
		maxDepth := 0

		for _, c := range group {
			if c == '(' {
				depth++
				if depth > maxDepth {
					maxDepth = depth
				}
			} else if c == ')' {
				depth--
			}
		}

		result = append(result, maxDepth)
	}

	return result
}

func main() {
	// Test the function with the example from docstring
	input := "(()()) ((())) () ((())()())"
	result := ParseNestedParens(input)
	fmt.Println(result)
	// Output: [2 3 1 3]
}