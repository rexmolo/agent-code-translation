package main

// Intersection determines if the length of intersection of two closed intervals is a prime number.
// Returns "YES" if the intersection length is prime, "NO" otherwise.
func Intersection(interval1 [2]int, interval2 [2]int) string {
	isPrime := func(num int) bool {
		if num == 1 || num == 0 {
			return false
		}
		if num == 2 {
			return true
		}
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	l := max(interval1[0], interval2[0])
	r := min(interval1[1], interval2[1])
	length := r - l

	if length > 0 && isPrime(length) {
		return "YES"
	}
	return "NO"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
