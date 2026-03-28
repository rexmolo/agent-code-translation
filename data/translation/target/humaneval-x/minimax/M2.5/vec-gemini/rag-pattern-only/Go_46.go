package main

func Fib4(n int) int {
	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[3] + results[2] + results[1] + results[0]
		results = append(results[1:], next)
	}

	return results[3]
}