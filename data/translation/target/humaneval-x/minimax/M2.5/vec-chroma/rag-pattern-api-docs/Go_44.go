package main

import (
	"fmt"
	"strconv"
)

func ChangeBase(x int, base int) string {
	ret := ""
	for x > 0 {
		ret = strconv.Itoa(x%base) + ret
		x /= base
	}
	return ret
}

func main() {
	fmt.Println(ChangeBase(8, 3)) // "22"
	fmt.Println(ChangeBase(8, 2)) // "1000"
	fmt.Println(ChangeBase(7, 2)) // "111"
	fmt.Println(ChangeBase(0, 2))  // "" (edge case)
}