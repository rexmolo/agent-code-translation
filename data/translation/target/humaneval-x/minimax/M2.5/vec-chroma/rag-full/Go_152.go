package main

import "fmt"

// abs returns the absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Compare takes two slices of equal length and returns a slice
// containing the absolute difference between each corresponding pair.
// If the guess matches exactly, the value is 0; otherwise, it's the
// absolute difference between guess and game.
func Compare(game, guess []int) []int {
	result := make([]int, 0, len(game))
	for i := 0; i < len(game); i++ {
		result = append(result, abs(game[i]-guess[i]))
	}
	return result
}

func main() {
	// Example usage:
	result1 := Compare([]int{1, 2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 2, -2})
	fmt.Println(result1) // [0 0 0 0 3 3]

	result2 := Compare([]int{0, 5, 0, 0, 0, 4}, []int{4, 1, 1, 0, 0, -2})
	fmt.Println(result2) // [4 4 1 0 0 6]
}
