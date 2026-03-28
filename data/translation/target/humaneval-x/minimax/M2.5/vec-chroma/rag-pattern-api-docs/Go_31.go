package main

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for k := 2; k < n; k++ {
		if n%k == 0 {
			return false
		}
	}
	return true
}
