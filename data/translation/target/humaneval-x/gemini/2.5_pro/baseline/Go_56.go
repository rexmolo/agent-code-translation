package main

import "fmt"

// CorrectBracketing checks if a string of brackets is correctly matched.
// It returns true if every opening bracket has a corresponding closing bracket.
func CorrectBracketing(brackets string) bool {
	depth := 0
	for _, b := range brackets {
		if b == '<' {
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

func main() {
	// Demonstration based on Python docstring examples
	testCases := []string{"<", "<>", "<<><>>", "><<>"}
	
	for _, tc := range testCases {
		fmt.Printf("CorrectBracketing(\"%s\") -> %v\n", tc, CorrectBracketing(tc))
	}
}