package main

// GenerateIntegers returns the even digits between a and b (inclusive),
// in ascending order. The result is bounded between 2 and 8.
func GenerateIntegers(a, b int) []int {
	// Normalize the range: find min and max of a, b
	minVal := a
	if b < minVal {
		minVal = b
	}
	maxVal := a
	if b > maxVal {
		maxVal = b
	}

	// Apply bounds: lower at least 2, upper at most 8
	lower := minVal
	if 2 > lower {
		lower = 2
	}
	upper := maxVal
	if 8 < upper {
		upper = 8
	}

	// Collect even numbers in the range
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
	println(GenerateIntegers(2, 8))   // [2 4 6 8]
	println(GenerateIntegers(8, 2))   // [2 4 6 8]
	println(GenerateIntegers(10, 14)) // []
}
