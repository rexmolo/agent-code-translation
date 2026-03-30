package main

import "fmt"

// SumProduct returns a fixed-size array of two integers
// representing the sum and product of all integers in the input slice.
// For an empty slice, sum is 0 and product is 1.
//
// Example:
//
//	SumProduct([]int{})      // returns [2]int{0, 1}
//	SumProduct([]int{1, 2, 3, 4})  // returns [2]int{10, 24}
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
	// Example usage
	fmt.Println(SumProduct([]int{}))
	fmt.Println(SumProduct([]int{1, 2, 3, 4}))
}
