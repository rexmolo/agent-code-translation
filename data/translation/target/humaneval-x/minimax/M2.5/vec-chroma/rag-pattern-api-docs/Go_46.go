package main

func Fib4(n int) int {
	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		// Slide window: remove first element, add new element
		results = append(results[1:], next)
	}

	return results[3]
}

func main() {
	// Example usage
	println(Fib4(5)) // Output: 4
	println(Fib4(6)) // Output: 8
	println(Fib4(7)) // Output: 14
}
