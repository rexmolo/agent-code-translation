package main

import "fmt"

func IsNested(s string) bool {
	// Collect indices of opening and closing brackets
	var openingIndices []int
	var closingIndices []int

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			openingIndices = append(openingIndices, i)
		} else {
			closingIndices = append(closingIndices, i)
		}
	}

	// Reverse the closing indices
	for i, j := 0, len(closingIndices)-1; i < j; i, j = i+1, j-1 {
		closingIndices[i], closingIndices[j] = closingIndices[j], closingIndices[i]
	}

	cnt := 0
	l := len(closingIndices)
	for i, idx := range openingIndices {
		if i < l && idx < closingIndices[i] {
			cnt++
		}
	}

	return cnt >= 2
}

func main() {
	// Test cases
	fmt.Println(IsNested("[[]]"))       // true
	fmt.Println(IsNested("[]]]]]]][[[[[]")) // false
	fmt.Println(IsNested("[][]"))       // false
	fmt.Println(IsNested("[]"))         // false
	fmt.Println(IsNested("[[][]]"))    // true
	fmt.Println(IsNested("[[]][["))    // true
}
