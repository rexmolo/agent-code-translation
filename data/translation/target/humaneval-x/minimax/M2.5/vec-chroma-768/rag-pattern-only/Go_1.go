package main

import (
	"fmt"
	"strings"
)

func SeparateParenGroups(parenString string) []string {
	result := []string{}
	currentString := strings.Builder{}
	currentDepth := 0

	for _, c := range parenString {
		if c == '(' {
			currentDepth++
			currentString.WriteRune(c)
		} else if c == ')' {
			currentDepth--
			currentString.WriteRune(c)

			if currentDepth == 0 {
				result = append(result, currentString.String())
				currentString.Reset()
			}
		}
		// Ignore spaces and other characters
	}

	return result
}

func main() {
	// Test the function
	input := "( ) (( )) (( )( ))"
	result := SeparateParenGroups(input)
	fmt.Println(result)
}