package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Multiply(a, b int) int {
	return abs(a%10) * abs(b%10)
}

func main() {
	fmt.Println(Multiply(148, 412))   // 16
	fmt.Println(Multiply(19, 28))     // 72
	fmt.Println(Multiply(2020, 1851)) // 0
	fmt.Println(Multiply(14, -15))    // 20
}
