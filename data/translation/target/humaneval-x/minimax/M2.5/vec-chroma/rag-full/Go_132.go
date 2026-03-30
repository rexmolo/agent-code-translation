package main

import "slices"

func IsNested(s string) bool {
	var openingBracketIndex []int
	var closingBracketIndex []int

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			openingBracketIndex = append(openingBracketIndex, i)
		} else {
			closingBracketIndex = append(closingBracketIndex, i)
		}
	}

	slices.Reverse(closingBracketIndex)

	cnt := 0
	i := 0
	l := len(closingBracketIndex)

	for _, idx := range openingBracketIndex {
		if i < l && idx < closingBracketIndex[i] {
			cnt++
			i++
		}
	}

	return cnt >= 2
}

func main() {
	// Test cases
	println(IsNested("[[]]"))      // true
	println(IsNested("[]]]]]]][[[[[]")) // false
	println(IsNested("[][]"))      // false
	println(IsNested("[]"))        // false
	println(IsNested("[[][]]")   // true
	println(IsNested("[[]][["))   // true
}
