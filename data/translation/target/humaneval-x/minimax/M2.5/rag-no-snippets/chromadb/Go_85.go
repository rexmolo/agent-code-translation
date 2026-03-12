package main

import "fmt"

func Add(lst []int) int {
	sum := 0
	for i := 1; i < len(lst); i += 2 {
		if lst[i]%2 == 0 {
			sum += lst[i]
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(Add([]int{4, 2, 6, 7})) // Expected: 2
	fmt.Println(Add([]int{1, 2, 3, 4, 5, 6})) // Expected: 2 + 4 + 6 = 12
}