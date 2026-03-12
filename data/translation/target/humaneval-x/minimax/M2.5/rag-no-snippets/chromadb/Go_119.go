package main

import "bufio"
import "fmt"
import "os"

func MatchParens(lst []string) string {
	// Helper function to check if a string has balanced parentheses
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
	scanner := bufio.NewScanner(os.Stdin)
	// Read first line (number of test cases, or just read the two strings)
	if !scanner.Scan() {
		return
	}
	// For simplicity, assume input contains the two strings on separate lines
	var lst []string
	for scanner.Scan() {
		lst = append(lst, scanner.Text())
		if len(lst) == 2 {
			break
		}
	}
	fmt.Println(MatchParens(lst))
}
