package main

import (
	"fmt"
)

// Eat calculates the total number of eaten carrots and the number of carrots remaining.
// It follows the logic from the original Python function.
// number: the number of carrots already eaten.
// need: the number of additional carrots to eat.
// remaining: the number of carrots available in stock.
// It returns a slice containing two integers: [total carrots eaten, carrots left].
func Eat(number, need, remaining int) []int {
	if need <= remaining {
		return []int{number + need, remaining - need}
	}
	// If there are not enough remaining carrots, eat all of them.
	return []int{number + remaining, 0}
}

// main function to test the Eat function with examples.
func main() {
	fmt.Println(Eat(5, 6, 10))
	fmt.Println(Eat(4, 8, 9))
	fmt.Println(Eat(1, 10, 10))
	fmt.Println(Eat(2, 11, 5))
}