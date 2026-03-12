package main

import (
	"fmt"
	"strings"
)

func ParseNestedParens(parenString string) []int {
	groups := strings.Split(parenString, " ")

	result := make([]int, 0)

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
			} else {
				depth--
			}
		}

		result = append(result, maxDepth)
	}

	return result
}

func main() {
	// Test the function
	result := ParseNestedParens("(()()) ((())) () ((())()())")
	fmt.Println(result)
	// Output: [2 3 1 3]
}
