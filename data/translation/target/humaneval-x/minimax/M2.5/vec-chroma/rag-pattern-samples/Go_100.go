package main

import "fmt"

func MakeAPile(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = n + 2*i
	}
	return result
}

func main() {
	// Test the function
	result := MakeAPile(3)
	fmt.Println(result)
}
