package main

import (
	"fmt"
	"slices"
)

func IsNested(s string) bool {
	openingBracketIndex := []int{}
	closingBracketIndex := []int{}

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
	fmt.Println(IsNested("[[]]"))           // Expected: true
	fmt.Println(IsNested("[]]]]]]][[[[[]")) // Expected: false
	fmt.Println(IsNested("[][]"))           // Expected: false
	fmt.Println(IsNested("[]"))             // Expected: false
	fmt.Println(IsNested("[[][]]"))          // Expected: true
	fmt.Println(IsNested("[[]][["))          // Expected: true
}
