func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for k := 2; k < n-1; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}
