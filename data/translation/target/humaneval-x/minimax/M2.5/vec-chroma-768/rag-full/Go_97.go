package main

import "fmt"

// Multiply returns the product of the unit digits of two integers
func Multiply(a, b int) int {
	return absInt(a%10) * absInt(b%10)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Multiply(148, 412))   // 16
	fmt.Println(Multiply(19, 28))     // 72
	fmt.Println(Multiply(2020, 1851)) // 0
	fmt.Println(Multiply(14, -15))    // 20
}
