func Digits(n int) int {
	product := 1
	oddCount := 0

	// Handle the case when n is 0
	if n == 0 {
		return 0
	}

	// Extract each digit from right to left using modulo
	for n > 0 {
		digit := n % 10
		if digit%2 == 1 {
			product *= digit
			oddCount++
		}
		n /= 10
	}

	if oddCount == 0 {
		return 0
	}
	return product
}
