package main

import "fmt"

func SumProduct(numbers []int) [2]int {
	sum := 0
	prod := 1

	for _, n := range numbers {
		sum += n
		prod *= n
	}
	return [2]int{sum, prod}
}

func main() {
	// Test cases from docstring
	fmt.Println(SumProduct([]int{}))
	fmt.Println(SumProduct([]int{1, 2, 3, 4}))
}