package main

import "strings"

func SeparateParenGroups(parenString string) []string {
	result := []string{}
	currentString := []rune{}
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
				currentString = []rune{}
			}
		}
	}

	return result
}