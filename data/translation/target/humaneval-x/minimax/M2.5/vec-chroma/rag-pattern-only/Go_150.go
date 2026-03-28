package main

import "fmt"

func XOrY(n, x, y int) int {
	if n <= 1 {
		return y
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return y
		}
	}
	return x
}

func main() {
	fmt.Println(XOrY(7, 34, 12)) // Output: 34
	fmt.Println(XOrY(15, 8, 5))  // Output: 5
}