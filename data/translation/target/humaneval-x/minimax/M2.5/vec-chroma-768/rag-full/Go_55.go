// Fib returns the n-th Fibonacci number.
func Fib(n int) int {
	if n == 0 {
		0
	}
	if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}