package main

import (
	"fmt"
	"strings"
)

func SeparateParenGroups(parenString string) []string {
	result := make([]string, 0)
	var currentString []byte
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
				currentString = currentString[:0]
			}
		}
		// Other characters (including spaces) are ignored
	}

	return result
}

func main() {
	// Test the function
	result := SeparateParenGroups("( ) (( )) (( )( ))")
	fmt.Println(result)

	// Test case: '()' should return ["()"]
	fmt.Println(SeparateParenGroups("()"))
	
	// Test case: '(())' should return ["(())"]
	fmt.Println(SeparateParenGroups("(())"))
}
