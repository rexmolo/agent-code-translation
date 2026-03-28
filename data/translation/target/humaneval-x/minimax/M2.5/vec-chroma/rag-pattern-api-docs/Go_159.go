package main

import "fmt"

func Eat(number, need, remaining int) []int {
	if need <= remaining {
		return []int{number + need, remaining - need}
	}
	return []int{number + remaining, 0}
}

func main() {
	// Test cases based on the examples
	fmt.Println(Eat(5, 6, 10))  // [11 4]
	fmt.Println(Eat(4, 8, 9))   // [12 1]
	fmt.Println(Eat(1, 10, 10)) // [11 0]
	fmt.Println(Eat(2, 11, 5))  // [7 0]
}
