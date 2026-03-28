package main

import "fmt"

func Multiply(a, b int) int {
	aLast := a % 10
	bLast := b % 10
	if aLast < 0 {
		aLast = -aLast
	}
	if bLast < 0 {
		bLast = -bLast
	}
	return aLast * bLast
}

func main() {
	fmt.Println(Multiply(148, 412)) // 16
	fmt.Println(Multiply(19, 28))  // 72
	fmt.Println(Multiply(2020, 1851)) // 0
	fmt.Println(Multiply(14, -15)) // 20
}
