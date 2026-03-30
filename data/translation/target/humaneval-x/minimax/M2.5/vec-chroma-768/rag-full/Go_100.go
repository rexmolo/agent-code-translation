package main

import "fmt"

func MakeAPile(n int) []int {
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, n+2*i)
	}
	return result
}

func main() {
	// Test with example: make_a_pile(3) should return [3, 5, 7]
	fmt.Println(MakeAPile(3))
}