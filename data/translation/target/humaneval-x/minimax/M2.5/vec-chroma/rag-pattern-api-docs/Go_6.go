package main

import (
	"strings"
)

func ParseNestedParens(parenString string) []int {
	// Split the input string by spaces
	groups := strings.Split(parenString, " ")

	var result []int

	// Process each group of parentheses
	for _, group := range groups {
		if group == "" {
			continue
		}

		// Calculate the maximum nesting depth for this group
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
	// Example usage
	input := "(()()) ((())) () ((())()')"
	result := ParseNestedParens(input)
	// Result: [2, 3, 1, 3]
	_ = result
}
