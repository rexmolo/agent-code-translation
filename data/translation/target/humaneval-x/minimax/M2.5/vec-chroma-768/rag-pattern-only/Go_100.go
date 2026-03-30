package main

import "fmt"

func MakeAPile(n int) []int {
	// Pre-allocate the slice with exact size needed
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = n + 2*i
	}
	return result
}

func main() {
	// Example usage
	result := MakeAPile(3)
	fmt.Println(result)
}