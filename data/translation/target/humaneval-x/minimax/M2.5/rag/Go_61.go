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
	fmt.Println(CorrectBracketing("("))      // Expected: false
	fmt.Println(CorrectBracketing("()"))     // Expected: true
	fmt.Println(CorrectBracketing("(()())")) // Expected: true
	fmt.Println(CorrectBracketing(")(()"))  // Expected: false
}
