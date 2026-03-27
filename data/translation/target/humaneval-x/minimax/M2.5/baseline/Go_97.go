package main

import "fmt"

func Multiply(a, b int) int {
	aUnit := a % 10
	if aUnit < 0 {
		aUnit = -aUnit
	}
	bUnit := b % 10
	if bUnit < 0 {
		bUnit = -bUnit
	}
	return aUnit * bUnit
}

func main() {
	// Test cases
	fmt.Println(Multiply(148, 412)) // 16
	fmt.Println(Multiply(19, 28))   // 72
	fmt.Println(Multiply(2020, 1851)) // 0
	fmt.Println(Multiply(14, -15))  // 20
}