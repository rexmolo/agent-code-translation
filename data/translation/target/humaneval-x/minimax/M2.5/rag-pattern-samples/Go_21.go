package main

func RescaleToUnit(numbers []float64) []float64 {
	// Find min and max values
	minNumber := numbers[0]
	maxNumber := numbers[0]
	for _, n := range numbers {
		if n < minNumber {
			minNumber = n
		}
		if n > maxNumber {
			maxNumber = n
		}
	}

	// Apply linear transform to scale values to [0, 1]
	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}
	return result
}