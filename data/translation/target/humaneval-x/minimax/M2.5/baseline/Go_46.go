package main

import "fmt"

func Fib4(n int) int {
	results := [4]int{0, 0, 2, 0}

	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		// Shift values: position 0 gets position 1, position 1 gets position 2, etc.
		results[0] = results[1]
		results[1] = results[2]
		results[2] = results[3]
		results[3] = next
	}

	return results[3]
}

func main() {
	fmt.Println(Fib4(5)) // Output: 4
	fmt.Println(Fib4(6)) // Output: 8
	fmt.Println(Fib4(7)) // Output: 14
}
