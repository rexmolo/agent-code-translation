package main

import (
	"fmt"
	"strconv"
)

func ChangeBase(x int, base int) string {
	if x == 0 {
		return "0"
	}
	ret := ""
	for x > 0 {
		ret = strconv.Itoa(x%base) + ret
		x = x / base
	}
	ret
}

func main() {
	// Test cases from Python docstring
	fmt.Println(ChangeBase(8, 3)) // '22'
	fmt.Println(ChangeBase(8, 2)) // '1000'
	fmt.Println(ChangeBase(7, 2)) // '111'
}
