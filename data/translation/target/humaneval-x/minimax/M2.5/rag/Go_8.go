package main

import (
	"fmt"
)

func SumProduct(numbers []int) [2]int {
	sumValue := 0
	prodValue := 1

	for _, n := range numbers {
		sumValue += n
		prodValue *= n
	}
	return [2]int{sumValue, prodValue}
}

func main() {
	// Test cases
	fmt.Println(SumProduct([]int{}))      // Output: [0 1]
	fmt.Println(SumProduct([]int{1, 2, 3, 4})) // Output: [10 24]
}
