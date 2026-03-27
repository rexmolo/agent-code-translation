package main

import "fmt"

func SeparateParenGroups(paren_string string) []string {
	result := []string{}
	currentString := []byte{}
	currentDepth := 0

	for _, c := range paren_string {
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
	// Example test
	result := SeparateParenGroups("( ) (( )) (( )( ))")
	fmt.Println(result)
}
