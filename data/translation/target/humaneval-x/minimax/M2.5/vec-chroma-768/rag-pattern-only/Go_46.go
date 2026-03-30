package main

func Fib4(n int) int {
	// Base cases for n < 4
	if n < 4 {
		switch n {
		case 0:
			return 0
		case 1:
			return 0
		case 2:
			return 2
		case 3:
			return 0
		}
	}

	// Sliding window using fixed 4-element array for efficiency
	results := [4]int{0, 0, 2, 0}

	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		// Shift values left: results[0] becomes results[1], etc.
		results[0] = results[1]
		results[1] = results[2]
		results[2] = results[3]
		results[3] = next
	}

	return results[3]
}

func main() {
	// Test cases
	println(Fib4(0)) // 0
	println(Fib4(5)) // 4
	println(Fib4(6)) // 8
	println(Fib4(7)) // 14
}