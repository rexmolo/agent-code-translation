func LargestPrimeFactor(n int) int {
	// Helper function to check if k is prime
	isPrime := func(k int) bool {
		if k < 2 {
			return false
		}
		for i := 2; i < k-1; i++ {
			if k%i == 0 {
				return false
			}
		}
		return true
	}

	largest := 1
	for j := 2; j <= n; j++ {
		if n%j == 0 && isPrime(j) {
			if j > largest {
				largest = j
			}
		}
	}
	return largest
}