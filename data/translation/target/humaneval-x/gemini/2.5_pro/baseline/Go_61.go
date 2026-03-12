package main

import (
	"fmt"
)

// CorrectBracketing checks if a string of brackets is correctly matched.
// brackets is a string of "(" and ")".
// It returns true if every opening bracket has a corresponding closing bracket.
func CorrectBracketing(brackets string) bool {
	depth := 0
	for _, b := range brackets {
		if b == '(' {
			depth++
		} else {
			depth--
		}
		if depth < 0 {
			return false
		}
	}
	return depth == 0
}

// main is the entry point of the program.
// It demonstrates the CorrectBracketing function with examples from the original Python doctests.
func main() {
	fmt.Println(CorrectBracketing("("))
	fmt.Println(CorrectBracketing("()"))
	fmt.Println(CorrectBracketing("(()())"))
	fmt.Println(CorrectBracketing(")(()"))
}
