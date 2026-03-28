package main

import (
	"fmt"
	"slices"
)

func RescaleToUnit(numbers []float64) []float64 {
	minNumber := slices.Min(numbers)
	maxNumber := slices.Max(numbers)

	result := make([]float64, len(numbers))
	for i, x := range numbers {
		result[i] = (x - minNumber) / (maxNumber - minNumber)
	}

	return result
}

func main() {
	numbers := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	result := RescaleToUnit(numbers)
	fmt.Println(result)
}
