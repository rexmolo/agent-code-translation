package main

import "fmt"

// MatchParens checks if concatenating two strings of parentheses in either
// order produces a balanced string.
func MatchParens(lst []string) string {
	check := func(s string) bool {
		val := 0
		for _, c := range s {
			if c == '(' {
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

	s1 := lst[0] + lst[1]
	s2 := lst[1] + lst[0]

	if check(s1) || check(s2) {
		return "Yes"
	}
	return "No"
}

func main() {
	// Test cases
	fmt.Println(MatchParens([]string{"()(", ")"})) // Yes
	fmt.Println(MatchParens([]string{")", ")"}))   // No
}
