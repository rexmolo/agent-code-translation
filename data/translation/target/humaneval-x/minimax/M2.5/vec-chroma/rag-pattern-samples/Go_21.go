package main

import "slices"

func RescaleToUnit(numbers []float64) []float64 {
    minNumber := slices.Min(numbers)
    maxNumber := slices.Max(numbers)
    
    rangeSize := maxNumber - minNumber
    
    result := make([]float64, len(numbers))
    for i, x := range numbers {
        result[i] = (x - minNumber) / rangeSize
    }
    return result
}