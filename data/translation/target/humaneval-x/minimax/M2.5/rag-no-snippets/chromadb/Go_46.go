func Fib4(n int) int {
	// Base cases: fib4(0) -> 0, fib4(1) -> 0, fib4(2) -> 2, fib4(3) -> 0
	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	// Iteratively compute fib4(n) using a sliding window of 4 elements
	for i := 4; i <= n; i++ {
		next := results[len(results)-1] + results[len(results)-2] + results[len(results)-3] + results[len(results)-4]
		results = append(results, next)
		results = results[1:] // Remove the oldest element (pop(0))
	}

	return results[len(results)-1]
}
