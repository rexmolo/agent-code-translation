package main

import "fmt"

func Multiply(a, b int) int {
	ua := a % 10
	ub := b % 10

	if ua < 0 {
		ua = -ua
	}
	if ub < 0 {
		ub = -ub
	}

	return ua * ub
}

func main() {
	fmt.Println(Multiply(148, 412))  // 16
	fmt.Println(Multiply(19, 28))    // 72
	fmt.Println(Multiply(2020, 1851))// 0
	fmt.Println(Multiply(14, -15))   // 20
}
