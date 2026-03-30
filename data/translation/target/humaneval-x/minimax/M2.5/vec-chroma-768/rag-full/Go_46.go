package main

func Fib4(n int) int {
	results := []int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[len(results)-1] + results[len(results)-2] + results[len(results)-3] + results[len(results)-4]
		results = append(results, next)
		results = results[1:] // pop(0) - remove the first element
	}

	return results[len(results)-1]
}
