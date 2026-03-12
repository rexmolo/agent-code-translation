package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	// Read input from stdin (format: two lines of space-separated integers)
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter game scores (space-separated):")
	scanner.Scan()
	gameLine := scanner.Text()
	gameStrs := strings.Fields(gameLine)
	game := make([]int, len(gameStrs))
	for i, s := range gameStrs {
		game[i], _ = strconv.Atoi(s)
	}
	
	fmt.Println("Enter guess scores (space-separated):")
	scanner.Scan()
	guessLine := scanner.Text()
	guessStrs := strings.Fields(guessLine)
	guess := make([]int, len(guessStrs))
	for i, s := range guessStrs {
		guess[i], _ = strconv.Atoi(s)
	}
	
	result := Compare(game, guess)
	
	// Output result as space-separated values
	for i, v := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
