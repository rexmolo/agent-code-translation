package main

// RescaleToUnit applies a linear transform to a slice of numbers,
// such that the smallest number becomes 0 and the largest becomes 1.
func RescaleToUnit(numbers []float64) []float64 {
	if (len(numbers) == 0) {
		return []float64{}
	}

	minNumber := numbers[0]
	maxNumber := numbers[0]
	for _, num := range numbers[1:] {
		if num < minNumber {
			minNumber = num
		}
		if num > maxNumber {
			maxNumber = num
		}
	}

	diff := maxNumber - minNumber
	result := make([]float64, len(numbers))

	// If min and max are the same, all elements are equal.
	// The range is 0. `make` initializes the slice with zeros, so we can return.
	if diff == 0 {
		return result
	}

	for i, x := range numbers {
		result[i] = (x - minNumber) / diff
	}

	return result
}
