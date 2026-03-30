package main

func Fib4(n int) int {
	// Initialize with base cases: fib4(0)=0, fib4(1)=0, fib4(2)=2, fib4(3)=0
	results := [4]int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		// Slide window: shift elements left and add new value at the end
		results[0] = results[1]
		results[1] = results[2]
		results[2] = results[3]
		results[3] = next
	}

	return results[3]
}
