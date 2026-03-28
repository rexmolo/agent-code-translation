package main

import (
	"fmt"
	"strings"
)

func SeparateParenGroups(parenString string) []string {
	result := make([]string, 0)
	var current strings.Builder
	currentDepth := 0

	for _, c := range parenString {
		switch c {
		case '(':
			currentDepth++
			current.WriteRune(c)
		case ')':
			currentDepth--
			current.WriteRune(c)

			if currentDepth == 0 {
				result = append(result, current.String())
				current.Reset()
			}
		}
	}

	return result
}

func main() {
	input := "( ) (( )) (( )( ))"
	result := SeparateParenGroups(input)
	fmt.Println(result)
}