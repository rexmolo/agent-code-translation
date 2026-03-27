package main

import "fmt"

func Digits(n int) int {
	product := 1
	oddCount := 0

	for n > 0 {
		digit := n % 10
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
		n /= 10
	}

	if oddCount == 0 {
		return 0
	}
	return product
}

func main() {
	// Test cases
	fmt.Println(Digits(1))   // == 1
	fmt.Println(Digits(4))   // == 0
	fmt.Println(Digits(235)) // == 15
}
