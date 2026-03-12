package main

import (
	"fmt"
)

func Fib4(n int) int {
	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		newVal := results[len(results)-1] + results[len(results)-2] + results[len(results)-3] + results[len(results)-4]
		results = append(results[1:], newVal)
	}

	return results[len(results)-1]
}

func main() {
	fmt.Println(Fib4(5))
	fmt.Println(Fib4(6))
	fmt.Println(Fib4(7))
}
