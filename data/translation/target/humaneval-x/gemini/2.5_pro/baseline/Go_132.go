package main

import "fmt"

// IsNested checks if a string of square brackets has a valid subsequence
// where at least one bracket is nested. The logic is a direct translation
// of the provided Python code.
func IsNested(s string) bool {
	var openingBracketIndex []int
	var closingBracketIndex []int

	for i, char := range s {
		if char == '[' {
			openingBracketIndex = append(openingBracketIndex, i)
		} else {
			closingBracketIndex = append(closingBracketIndex, i)
		}
	}

	// Reverse the closingBracketIndex slice in-place
	for i, j := 0, len(closingBracketIndex)-1; i < j; i, j = i+1, j-1 {
		closingBracketIndex[i], closingBracketIndex[j] = closingBracketIndex[j], closingBracketIndex[i]
	}

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
	fmt.Printf("[[]] -> %t\n", IsNested("[[]]"))
	fmt.Printf("[]]]]]]][[[[[] -> %t\n", IsNested("[]]]]]]][[[[[]"))
	fmt.Printf("[][] -> %t\n", IsNested("[][]"))
	fmt.Printf("[] -> %t\n", IsNested("[]"))
	fmt.Printf("[[][]] -> %t\n", IsNested("[[][]]"))
	fmt.Printf("[[]][[ -> %t\n", IsNested("[[]][["))
}
