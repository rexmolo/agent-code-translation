package main

import (
	"fmt"
	"math"
)

// Compare determines if a person correctly guessed the results of a number of matches.
// It's given two slices of scores and guesses of equal length, where each index shows a match.
// It returns a slice of the same length denoting how far off each guess was. If they have guessed correctly,
// the value is 0, and if not, the value is the absolute difference between the guess and the score.
func Compare(game, guess []int) []int {
	// The problem statement guarantees game and guess are of equal length.
	results := make([]int, len(game))

	for i, gameScore := range game {
		guessScore := guess[i]
		// Go's math.Abs function works on float64, so we cast, calculate, and cast back.
		// An alternative without casting would be a simple if check:
		// diff := gameScore - guessScore
		// if diff < 0 {
		// 	diff = -diff
		// }
		// results[i] = diff
		diff := float64(gameScore - guessScore)
		results[i] = int(math.Abs(diff))
	}

	return results
}

func main() {
	// Example 1:
	game1 := []int{1, 2, 3, 4, 5, 1}
	guess1 := []int{1, 2, 3, 4, 2, -2}
	result1 := Compare(game1, guess1)
	fmt.Printf("compare(%v, %v) -> %v\n", game1, guess1, result1)

	// Example 2:
	game2 := []int{0, 5, 0, 0, 0, 4}
	guess2 := []int{4, 1, 1, 0, 0, -2}
	result2 := Compare(game2, guess2)
	fmt.Printf("compare(%v, %v) -> %v\n", game2, guess2, result2)
}
