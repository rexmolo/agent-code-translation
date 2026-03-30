package main

import "math"

func Compare(game, guess []int) []int {
	result := make([]int, len(game))
	for i := range game {
		diff := game[i] - guess[i]
		if diff < 0 {
			diff = -diff
		}
		result[i] = diff
	}
	return result
}

func main() {
	// Example usage
	game := []int{1, 2, 3, 4, 5, 1}
	guess := []int{1, 2, 3, 4, 2, -2}
	result := Compare(game, guess)
	// result should be [0, 0, 0, 0, 3, 3]
}
