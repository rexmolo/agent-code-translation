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
	fmt.Println(CorrectBracketing("<>"))      // true
	fmt.Println(CorrectBracketing("<<><>>"))  // true
	fmt.Println(CorrectBracketing("><<>"))    // false
	fmt.Println(CorrectBracketing("<>><"))    // false
}