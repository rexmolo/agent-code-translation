package main

import "fmt"

func Derivative(xs []int) []int {
	result := make([]int, 0, len(xs)-1)
	for i, x := range xs {
		if i > 0 {
			result = append(result, i*x)
		}
	}
	return result
}

func main() {
	fmt.Println(Derivative([]int{3, 1, 2, 4, 5}))
	fmt.Println(Derivative([]int{1, 2, 3}))
}
