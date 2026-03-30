package main

import "fmt"

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
	// Test cases
	tests := []string{"<>", "<<><>>", "<>", "><<>", ""}
	expected := []bool{false, true, true, false, true}
	
	for i, test := range tests {
		result := CorrectBracketing(test)
		fmt.Printf("CorrectBracketing(%q) = %v (expected %v)\n", test, result, expected[i])
	}
}
