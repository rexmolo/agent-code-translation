package main

import (
    "math"
)

func SumSquares(lst []float64) int {
    squared := 0
    for _, i := range lst {
        squared += int(math.Ceil(i) * math.Ceil(i))
    }
    return squared
}
