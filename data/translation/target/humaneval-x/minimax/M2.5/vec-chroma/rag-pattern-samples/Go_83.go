package main

import "fmt"

func StartsOneEnds(n int) int {
	if n == 1 {
		return 1
	}

	// 18 * 10^(n-2)
	// Using integer multiplication to avoid float precision issues
	result := 18
	for i := 0; i < n-2; i++ {
		result *= 10
	}
	return result
}

func main() {
	// Test the function
	fmt.Println(StartsOneEnds(1)) // 1
	fmt.Println(StartsOneEnds(2)) // 18
	fmt.Println(StartsOneEnds(3)) // 180
}