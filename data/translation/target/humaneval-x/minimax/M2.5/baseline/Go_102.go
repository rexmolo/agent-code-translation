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
	// Test cases
	fmt.Println(ChooseNum(12, 15)) // Expected: 14
	fmt.Println(ChooseNum(13, 12)) // Expected: -1
}
