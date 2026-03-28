package main

import "fmt"

// GenerateIntegers returns the even digits between a and b (inclusive),
// in ascending order. The range is constrained to [2, 8].
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
	// Example usage
	fmt.Println(GenerateIntegers(2, 8))  // [2 4 6 8]
	fmt.Println(GenerateIntegers(8, 2))  // [2 4 6 8]
	fmt.Println(GenerateIntegers(10, 14)) // []
}
