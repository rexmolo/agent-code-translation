package main

import (
	"fmt"
)

func Derivative(xs []int) []int {
	if len(xs) <= 1 {
		return []int{}
	}
	
	result := make([]int, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result[i-1] = i * x
		}
	}
	return result
}

func main() {
	fmt.Println(Derivative([]int{3, 1, 2, 4, 5})) // [1 4 12 20]
	fmt.Println(Derivative([]int{1, 2, 3}))       // [2 6]
}
