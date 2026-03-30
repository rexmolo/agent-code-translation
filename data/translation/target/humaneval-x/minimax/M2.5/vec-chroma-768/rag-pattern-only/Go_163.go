package main

// GenerateIntegers returns even digits between two positive integers a and b,
// in ascending order. The range is clamped between 2 and 8.
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
	// Example usage - reading from stdin and writing to stdout
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}
		result := GenerateIntegers(a, b)
		fmt.Println(result)
	}
}
