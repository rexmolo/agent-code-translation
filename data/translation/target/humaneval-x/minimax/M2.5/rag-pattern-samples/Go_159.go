package main

import "fmt"

func Eat(number, need, remaining int) []int {
	if need <= remaining {
		return []int{number + need, remaining - need}
	}
	return []int{number + remaining, 0}
}

func main() {
	// Example test cases
	result1 := Eat(5, 6, 10)
	fmt.Println(result1)
	
	result2 := Eat(4, 8, 9)
	fmt.Println(result2)
	
	result3 := Eat(1, 10, 10)
	fmt.Println(result3)
	
	result4 := Eat(2, 11, 5)
	fmt.Println(result4)
}