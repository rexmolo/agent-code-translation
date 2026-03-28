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
	fmt.Println(ChangeBase(8, 3)) // Output: 22
	fmt.Println(ChangeBase(8, 2)) // Output: 1000
	fmt.Println(ChangeBase(7, 2)) // Output: 111
}
