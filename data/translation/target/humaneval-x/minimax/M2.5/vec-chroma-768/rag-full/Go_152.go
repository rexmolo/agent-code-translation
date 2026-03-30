package main

// Compare takes two slices of equal length and returns a slice
// where each element is the absolute difference between the
// corresponding elements of the input slices.
func Compare(game, guess []int) []int {
	result := make([]int, len(game))
	for i := 0; i < len(game); i++ {
		diff := game[i] - guess[i]
		if diff < 0 {
			diff = -diff
		}
		result[i] = diff
	}
	return result
}
