package main

import "fmt"

func ChangeBase(x int, base int) string {
	if x == 0 {
		return "0"
	}
	ret := ""
	for x > 0 {
		ret = string(rune('0'+x%base)) + ret
		x /= base
	}
	ret
}

func main() {
	fmt.Println(ChangeBase(8, 3))
	fmt.Println(ChangeBase(8, 2))
	fmt.Println(ChangeBase(7, 2))
}
