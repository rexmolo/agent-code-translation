package main

func GenerateIntegers(a, b int) []int {
	// lower = max(2, min(a, b))
	lower := a
	if b < lower {
		lower = b
	}
	if lower < 2 {
		lower = 2
	}

	// upper = min(8, max(a, b))
	upper := a
	if b > upper {
		upper = b
	}
	if upper > 8 {
		upper = 8
	}

	// Return even numbers from lower to upper
	var result []int
	for i := lower; i <= upper; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}