package main

import "fmt"

func ChooseNum(x, y int) int {
	if x > y {
		return -1
	}
	if y%2 == 0 {
		return y
	}
	if x == y {
		return -1
	}
	return y - 1
}

func main() {
	fmt.Println(ChooseNum(12, 15)) // Output: 14
	fmt.Println(ChooseNum(13, 12)) // Output: -1
}