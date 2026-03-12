package main

import "fmt"

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
	result1 := SumProduct([]int{})
	fmt.Printf("%v\n", result1)

	result2 := SumProduct([]int{1, 2, 3, 4})
	fmt.Printf("%v\n", result2)
}
