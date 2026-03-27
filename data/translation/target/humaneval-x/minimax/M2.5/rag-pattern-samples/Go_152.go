package main

// abs returns the absolute value of an integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Compare takes two slices of equal length and returns a slice
// where each element is the absolute difference between the
// corresponding elements of game and guess.
func Compare(game, guess []int) []int {
	result := make([]int, 0, len(game))
	for i := 0; i < len(game); i++ {
		result = append(result, abs(game[i]-guess[i]))
	}
	return result
}
