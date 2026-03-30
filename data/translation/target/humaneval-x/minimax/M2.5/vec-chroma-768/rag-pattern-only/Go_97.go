package main

import "math"

func Multiply(a, b int) int {
	return abs(a%10) * abs(b%10)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Test using math.Abs
func Multiply2(a, b int) int {
	return int(math.Abs(float64(a%10))) * int(math.Abs(float64(b%10)))
}

func main() {
	// Example tests (you can add more)
	println(Multiply(148, 412))  // 16
	println(Multiply(19, 28))     // 72
	println(Multiply(2020, 1851)) // 0
	println(Multiply(14, -15))    // 20
}
