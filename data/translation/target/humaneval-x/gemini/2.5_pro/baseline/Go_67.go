package main

import (
	"fmt"
	"strconv"
	"strings"
)

// FruitDistribution calculates the number of mangoes in a basket.
// It takes a string `s` describing the number of apples and oranges,
// and an integer `n` representing the total number of fruits.
func FruitDistribution(s string, n int) int {
	// Split the input string by spaces to get individual words.
	words := strings.Split(s, " ")
	
	// Initialize a variable to hold the sum of apples and oranges.
	applesAndOrangesCount := 0

	// Iterate over each word.
	for _, word := range words {
		// Try to convert the word to an integer.
		num, err := strconv.Atoi(word)
		
		// If the conversion is successful (error is nil), the word is a number.
		if err == nil {
			// Add the number to our running total.
			applesAndOrangesCount += num
		}
	}

	// The number of mangoes is the total number of fruits minus the sum of apples and oranges.
	return n - applesAndOrangesCount
}

// main function to test the FruitDistribution function with examples.
func main() {
	fmt.Println(FruitDistribution("5 apples and 6 oranges", 19))
	fmt.Println(FruitDistribution("0 apples and 1 oranges", 3))
	fmt.Println(FruitDistribution("2 apples and 3 oranges", 100))
	fmt.Println(FruitDistribution("100 apples and 1 oranges", 120))
}
