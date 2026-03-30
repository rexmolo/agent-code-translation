package main

import (
	"fmt"
	"strings"
)

func SeparateParenGroups(parenString string) []string {
	result := []string{}
	currentString := []rune{}
	currentDepth := 0

	for _, c := range parenString {
		if c == ' ' {
			continue // Ignore spaces
		}
		if c == '(' {
			currentDepth++
			currentString = append(currentString, c)
		} else if c == ')' {
			currentDepth--
			currentString = append(currentString, c)

			if currentDepth == 0 {
				result = append(result, string(currentString))
				currentString = nil // Clear the slice
			}
		}
	}

	return result
}

func main() {
	// Test the function
	input := "( ) (( )) (( )( ))"
	result := SeparateParenGroups(input)
	fmt.Println(result)
	// Output: [() (()) (()())]
}
