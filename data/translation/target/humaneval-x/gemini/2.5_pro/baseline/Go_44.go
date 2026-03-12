package main

import (
	"fmt"
	"strconv"
)

// ChangeBase changes the numerical base of input number x to a new base.
// It returns the string representation of the number in the new base.
// This implementation assumes the base is less than 10.
func ChangeBase(x int, base int) string {
	if x == 0 {
		return "0"
	}
	var res string
	for x > 0 {
		remainder := x % base
		res = strconv.Itoa(remainder) + res
		x /= base
	}
	return res
}

// main function to demonstrate ChangeBase with examples from the original docstring.
func main() {
	fmt.Println(ChangeBase(8, 3))
	fmt.Println(ChangeBase(8, 2))
	fmt.Println(ChangeBase(7, 2))
}
