package main

import "strings"

// SeparateParenGroups separates a string containing multiple groups of nested parentheses
// into a slice of strings.
//
// This function mirrors the logic of the original Python version.
// It identifies top-level balanced parenthesis groups and extracts them.
// Characters that are not parentheses are ignored.
func SeparateParenGroups(paren_string string) []string {
	var result []string
	var currentString strings.Builder
	currentDepth := 0

	for _, r := range paren_string {
		if r == '(' {
			currentDepth++
			currentString.WriteRune(r)
		} else if r == ')' {
			currentDepth--
			currentString.WriteRune(r)

			if currentDepth == 0 {
				result = append(result, currentString.String())
				currentString.Reset()
			}
		}
	}

	return result
}
