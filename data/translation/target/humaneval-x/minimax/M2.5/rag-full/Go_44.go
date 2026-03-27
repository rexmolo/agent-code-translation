package main

import (
	"fmt"
	"strconv"
)

// ChangeBase converts number x to the specified base and returns its string representation.
// Base must be less than 10.
// Examples:
// - ChangeBase(8, 3) returns "22"
// - ChangeBase(8, 2) returns "1000"
// - ChangeBase(7, 2) returns "111"
func ChangeBase(x int, base int) string {
	ret := ""
	for x > 0 {
		ret = strconv.Itoa(x%base) + ret
		x = x / base
	}
	return ret
}

func main() {
	// Test cases from docstring
	fmt.Println(ChangeBase(8, 3)) // Output: 22
	fmt.Println(ChangeBase(8, 2)) // Output: 1000
	fmt.Println(ChangeBase(7, 2)) // Output: 111
}