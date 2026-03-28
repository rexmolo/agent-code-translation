package main

import "fmt"

func Fib4(n int) int {
	// Base cases: fib4(0)=0, fib4(1)=0, fib4(2)=2, fib4(3)=0
	results := [4]int{0, 0, 2, 0}

	if n < 4 {
		return results[n]
	}

	// Iteratively compute the sequence using a sliding window of 4 elements
	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		// Shift elements to the left and add the new value at the end
		results[0] = results[1]
		results[1] = results[2]
		results[2] = results[3]
		results[3] = next
	}

	return results[3]
}

func main() {
	fmt.Println(Fib4(5)) // 4
	fmt.Println(Fib4(6)) // 8
	fmt.Println(Fib4(7)) // 14
}