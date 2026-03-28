package main

import "fmt"

func IsNested(s string) bool {
	var openingIndices []int
	var closingIndices []int

	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			openingIndices = append(openingIndices, i)
		} else {
			closingIndices = append(closingIndices, i)
		}
	}

	cnt := 0
	i := 0
	l := len(closingIndices)

	// Iterate through opening brackets and check against reversed closing indices
	for _, idx := range openingIndices {
		// Get closing bracket from the end (simulating reversal)
		if i < l && idx < closingIndices[l-1-i] {
			cnt++
			i++
		}
	}

	return cnt >= 2
}

func main() {
	fmt.Println(IsNested("[[]]"))       // true
	fmt.Println(IsNested("[]]]]]]][[[[[]")) // false
	fmt.Println(IsNested("[][]"))       // false
	fmt.Println(IsNested("[]"))          // false
	fmt.Println(IsNested("[[][]]"))     // true
	fmt.Println(IsNested("[[]][["))     // true
}
