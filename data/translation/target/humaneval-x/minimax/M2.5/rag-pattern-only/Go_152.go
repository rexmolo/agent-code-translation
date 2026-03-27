package main

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