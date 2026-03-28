package main

import "math"

func Compare(game, guess []int) []int {
    result := make([]int, 0, len(game))
    for i := 0; i < len(game); i++ {
        diff := game[i] - guess[i]
        // Convert to float64 for math.Abs, then back to int
        absDiff := int(math.Abs(float64(diff)))
        result = append(result, absDiff)
    }
    return result
}