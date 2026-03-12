func IsSimplePower(x int, n int) bool {
	if n == 1 {
		return x == 1
	}

	power := 1
	for power < x {
		// Check for potential overflow before multiplying
		if power > x/n {
			return false
		}
		power = power * n
	}
	return power == x
}