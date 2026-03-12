package main

func GenerateIntegers(a, b int) []int {
	// Given two positive integers a and b, return the even digits between a
	// and b, in ascending order.
	//
	// For example:
	// GenerateIntegers(2, 8) => [2, 4, 6, 8]
	// GenerateIntegers(8, 2) => [2, 4, 6, 8]
	// GenerateIntegers(10, 14) => []

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
