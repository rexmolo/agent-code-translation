package main

import "fmt"

// CorrectBracketing checks if every opening bracket "<" has a corresponding closing bracket ">".
// Returns true if the brackets are correctly matched, false otherwise.
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
	// Test cases matching the Python docstring examples
	fmt.Println(CorrectBracketing("<>"))    // Expected: false
	fmt.Println(CorrectBracketing("<>"))   // Expected: true
	fmt.Println(CorrectBracketing("<<><>>")) // Expected: true
	fmt.Println(CorrectBracketing("><<>")) // Expected: false
}
