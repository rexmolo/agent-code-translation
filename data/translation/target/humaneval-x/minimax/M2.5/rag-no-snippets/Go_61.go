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
	testCases := []string{"(", "()", "(()())", ")(()"}
	for _, tc := range testCases {
		fmt.Printf("%s -> %t\n", tc, CorrectBracketing(tc))
	}
}