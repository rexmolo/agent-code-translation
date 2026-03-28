package main

import (
	"fmt"
	"strings"
)

func ParseNestedParens(paren_string string) []int {
	// Helper function to parse a single paren group
	parseParenGroup := func(s string) int {
		depth := 0
		maxDepth := 0
		for _, c := range s {
			if c == '(' {
				depth++
				if depth > maxDepth {
					maxDepth = depth
				}
			} else if c == ')' {
				depth--
			}
		}
		return maxDepth
	}

	// Split by space and filter empty strings, then process each group
	groups := strings.Split(paren_string, " ")
	result := make([]int, 0, len(groups))

	for _, group := range groups {
		if group != "" {
			result = append(result, parseParenGroup(group))
		}
	}

	return result
}

func main() {
	// Test the function
	result := ParseNestedParens("(()()) ((())) () ((())()())")
	fmt.Println(result)
}
