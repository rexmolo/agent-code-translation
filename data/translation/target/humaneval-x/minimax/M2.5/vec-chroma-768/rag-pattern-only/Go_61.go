package main

import "fmt"

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

func main() {
	// Test cases
	testCases := []string{"(", "()", "(()())", ")(()"}
	for _, tc := range testCases {
		fmt.Printf("CorrectBracketing(%q) = %v\n", tc, CorrectBracketing(tc))
	}
}
