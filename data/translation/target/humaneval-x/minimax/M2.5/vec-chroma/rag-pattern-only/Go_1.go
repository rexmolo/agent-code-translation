package main

import "fmt"

func SeparateParenGroups(paren_string string) []string {
	result := []string{}
	var currentString []rune
	currentDepth := 0

	for _, c := range paren_string {
		switch c {
		case '(':
			currentDepth++
			currentString = append(currentString, c)
		case ')':
			currentDepth--
			currentString = append(currentString, c)
			if currentDepth == 0 {
				result = append(result, string(currentString))
				currentString = currentString[:0]
			}
		}
		// Spaces are ignored as they don't match any case
	}

	return result
}

func main() {
	// Test the function
	result := SeparateParenGroups("( ) (( )) (( )( ))")
	fmt.Println(result)
}