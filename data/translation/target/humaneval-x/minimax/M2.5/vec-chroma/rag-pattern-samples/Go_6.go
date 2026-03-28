package main

import (
	"fmt"
	"strings"
)

func ParseNestedParens(parenString string) []int {
	var result []int

	groups := strings.Split(parenString, " ")

	for _, group := range groups {
		if group == "" {
			continue
		}
		result = append(result, parseParenGroup(group))
	}

	return result
}

func parseParenGroup(s string) int {
	depth := 0
	maxDepth := 0

	for _, c := range s {
		if c == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else {
			depth--
		}
	}

	return maxDepth
}

func main() {
	// Test the function
	result := ParseNestedParens("(()()) ((())) () ((())()())")
	fmt.Println(result)
}