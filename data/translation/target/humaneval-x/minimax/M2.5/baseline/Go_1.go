package main

import (
	"strings"
)

func SeparateParenGroups(parenString string) []string {
	result := []string{}
	currentString := []rune{}
	currentDepth := 0

	for _, c := range parenString {
		if c == '(' {
			currentDepth++
			currentString = append(currentString, c)
		} else if c == ')' {
			currentDepth--
			currentString = append(currentString, c)

			if currentDepth == 0 {
				result = append(result, string(currentString))
				currentString = currentString[:0] // Clear the slice
			}
		}
		// Spaces are automatically ignored since we only process '(' and ')'
	}

	return result
}

func main() {
	// Example usage
	result := SeparateParenGroups("( ) (( )) (( )( ))")
	// Result: ["()", "(())", "(()())"]
	_ = result
}
