func Fibfib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	// Iterative approach - O(n) time, O(1) space
	// Using sliding window technique
	fib0, fib1, fib2 := 0, 0, 1 // fibfib(0), fibfib(1), fibfib(2)

	for i := 3; i <= n; i++ {
		fib0, fib1, fib2 = fib1, fib2, fib0+fib1+fib2
	}

	return fib2
}
