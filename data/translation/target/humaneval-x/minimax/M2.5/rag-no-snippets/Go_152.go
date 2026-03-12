package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Compare(game, guess []int) []int {
	result := make([]int, len(game))
	for i := 0; i < len(game); i++ {
		result[i] = abs(game[i] - guess[i])
	}
	return result
}

func main() {
	// Example tests
	fmt.Println(Compare([]int{1, 2, 3, 4, 5, 1}, []int{1, 2, 3, 4, 2, -2})) // [0 0 0 0 3 3]
	fmt.Println(Compare([]int{0, 5, 0, 0, 0, 4}, []int{4, 1, 1, 0, 0, -2})) // [4 4 1 0 0 6]
}
