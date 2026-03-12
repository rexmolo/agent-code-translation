package main

import (
	"fmt"
)

// check determines if a string of parentheses is balanced.
// It returns true if the string is balanced, and false otherwise.
func check(s string) bool {
	val := 0
	for _, char := range s {
		if char == '(' {
			val++
		} else {
			val--
		}
		if val < 0 {
			return false
		}
	}
	return val == 0
}

// MatchParens checks if it is possible to concatenate two strings of parentheses
// in some order to form a well-formed (balanced) string.
// It returns "Yes" if possible, and "No" otherwise.
func MatchParens(lst []string) string {
	// Concatenate in both possible orders
	s1 := lst[0] + lst[1]
	s2 := lst[1] + lst[0]

	// Check if either of the concatenated strings is balanced
	if check(s1) || check(s2) {
		return "Yes"
	}
	return "No"
}

func main() {
	// Example 1
	ex1 := []string{"()(", ")"}
	fmt.Printf("match_parens(%v) == '%s'\n", ex1, MatchParens(ex1)) // Expected: Yes

	// Example 2
	ex2 := []string{")", ")"}
	fmt.Printf("match_parens(%v) == '%s'\n", ex2, MatchParens(ex2)) // Expected: No

	// Example 3
	ex3 := []string{")(", "()"}
	fmt.Printf("match_parens(%v) == '%s'\n", ex3, MatchParens(ex3)) // Expected: Yes
}
