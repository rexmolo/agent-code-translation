func Fib(n int) int {
	// Return n-th Fibonacci number.
	// >>> Fib(10)
	// 55
	// >>> Fib(1)
	// 1
	// >>> Fib(8)
	// 21
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}
