package main

import "fmt"

func SeparateParenGroups(parenString string) []string {
	result := make([]string, 0)
	currentString := make([]rune, 0)
	currentDepth := 0

	for _, c := range parenString {
		if c == ' ' {
			continue
		}
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
	}

	return result
}

func main() {
	// Test the function
	result := SeparateParenGroups("( ) (( )) (( )( ))")
	fmt.Println(result)
}