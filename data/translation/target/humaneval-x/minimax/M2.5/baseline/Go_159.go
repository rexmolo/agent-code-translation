package main

import "fmt"

// Eat calculates the number of carrots eaten and remaining after a meal.
// It takes three integer parameters:
// - number: the number of carrots already eaten
// - need: the number of carrots the rabbit needs to eat
// - remaining: the number of carrots available in stock
// Returns a slice of two integers: [total eaten, remaining carrots]
func Eat(number, need, remaining int) []int {
	if need <= remaining {
		return []int{number + need, remaining - need}
	}
	return []int{number + remaining, 0}
}

func main() {
	// Example test cases
	result1 := Eat(5, 6, 10)
	fmt.Printf("eat(5, 6, 10) -> %v\n", result1)

	result2 := Eat(4, 8, 9)
	fmt.Printf("eat(4, 8, 9) -> %v\n", result2)

	result3 := Eat(1, 10, 10)
	fmt.Printf("eat(1, 10, 10) -> %v\n", result3)

	result4 := Eat(2, 11, 5)
	fmt.Printf("eat(2, 11, 5) -> %v\n", result4)
}
