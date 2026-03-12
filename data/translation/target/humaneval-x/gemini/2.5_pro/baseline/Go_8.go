package main

import "fmt"

// SumProduct calculates the sum and product of all integers in a slice.
// For an empty slice, the sum is 0 and the product is 1.
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
	// Example 1: Empty slice
	emptyCase := []int{}
	emptyResult := SumProduct(emptyCase)
	fmt.Printf("Input: %v -> Sum: %d, Product: %d\n", emptyCase, emptyResult[0], emptyResult[1])

	// Example 2: Non-empty slice
	numbers := []int{1, 2, 3, 4}
	result := SumProduct(numbers)
	fmt.Printf("Input: %v -> Sum: %d, Product: %d\n", numbers, result[0], result[1])
}
