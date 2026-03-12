func Fib4(n int) int {
	results := [4]int{0, 0, 2, 0}
	if n < 4 {
		return results[n]
	}

	for i := 4; i <= n; i++ {
		next := results[0] + results[1] + results[2] + results[3]
		results[0] = results[1]
		results[1] = results[2]
		results[2] = results[3]
		results[3] = next
	}

	return results[3]
}
