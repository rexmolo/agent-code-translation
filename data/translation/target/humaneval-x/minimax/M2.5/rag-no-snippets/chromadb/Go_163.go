package main

import (
	"fmt"
	"slices"
)

func GenerateIntegers(a, b int) []int {
	lower := max(2, min(a, b))
	upper := min(8, max(a, b))

	var result []int
	for i := lower; i <= upper; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	fmt.Println(GenerateIntegers(2, 8))
	fmt.Println(GenerateIntegers(8, 2))
	fmt.Println(GenerateIntegers(10, 14))
}