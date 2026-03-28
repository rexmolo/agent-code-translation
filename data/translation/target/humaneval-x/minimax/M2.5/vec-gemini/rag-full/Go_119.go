package main

import (
	"fmt"
)

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

	S1 := lst[0] + lst[1]
	S2 := lst[1] + lst[0]

	if check(S1) || check(S2) {
		return "Yes"
	}
	return "No"
}

func main() {
	// Example tests
	fmt.Println(MatchParens([]string{"()(", ")"}))    // Expected: Yes
	fmt.Println(MatchParens([]string{")", ")"}))      // Expected: No
}