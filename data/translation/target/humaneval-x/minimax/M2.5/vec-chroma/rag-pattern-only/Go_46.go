package main

func Fib4(n int) int {
	// Base cases: fib4(0) -> 0, fib4(1) -> 0, fib4(2) -> 2, fib4(3) -> 0
	results := [4]int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	// Use modular arithmetic to keep track of last 4 values
	for i := 4; i <= n; i++ {
		next := results[(i-1)%4] + results[(i-2)%4] + results[(i-3)%4] + results[(i-4)%4]
		results[i%4] = next
	}

	return results[n%4]
}

func main() {
	// Test cases
	println(Fib4(5)) // Expected: 4
	println(Fib4(6)) // Expected: 8
	println(Fib4(7)) // Expected: 14
}
