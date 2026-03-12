package main

import (
	"bufio"
	"fmt"
	"os"
)

func MatchParens(lst []string) string {
	// Helper function to check if string has balanced parentheses
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
	// Read input
	scanner := bufio.NewScanner(os.Stdin)
	var lst []string
	for scanner.Scan() {
		lst = append(lst, scanner.Text())
		if len(lst) == 2 {
			break
		}
	}
	fmt.Println(MatchParens(lst))
}
