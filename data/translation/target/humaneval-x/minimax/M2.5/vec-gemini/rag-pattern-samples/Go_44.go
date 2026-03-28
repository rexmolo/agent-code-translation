package main

import "fmt"

func ChangeBase(x int, base int) string {
	if x == 0 {
		return "0"
	}
	ret := ""
	for x > 0 {
		ret = fmt.Sprintf("%d%s", x%base, ret)
		x /= base
	}
	return ret
}

func main() {
	// Test cases from docstring
	fmt.Println(ChangeBase(8, 3)) // '22'
	fmt.Println(ChangeBase(8, 2)) // '1000'
	fmt.Println(ChangeBase(7, 2)) // '111'
}