package main

import "fmt"

func GenerateIntegers(a, b int) []int {
	lower := min(a, b)
	if lower < 2 {
		lower = 2
	}

	upper := max(a, b)
	if upper > 8 {
		upper = 8
	}

	result := make([]int, 0)
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
