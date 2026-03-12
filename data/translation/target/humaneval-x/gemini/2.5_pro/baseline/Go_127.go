package main

import "fmt"

// isPrime checks if a number is a prime number.
// The logic is a direct translation from the provided Python code.
func isPrime(num int) bool {
	if num <= 1 { // In Python: `if num == 1 or num == 0: return False`
		return false
	}
	// The loop starts at 2. If num is 2, the loop condition `i < num` is false,
	// and the function correctly returns true.
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Intersection determines if the length of the intersection of two intervals is a prime number.
func Intersection(interval1 [2]int, interval2 [2]int) string {
	// Find the start of the intersection, which is the maximum of the two start points.
	l := max(interval1[0], interval2[0])
	// Find the end of the intersection, which is the minimum of the two end points.
	r := min(interval1[1], interval2[1])

	// Calculate the length of the intersection as defined in the Python code.
	length := r - l

	// The Python code returns "YES" only if length > 0.
	// This implicitly handles cases of no intersection (l > r, so length is negative)
	// and single-point intersections (l = r, so length is 0).
	if length > 0 && isPrime(length) {
		return "YES"
	}

	return "NO"
}

// main function to test the implementation with examples from the problem description.
func main() {
	fmt.Printf("intersection((1, 2), (2, 3)) ==> \"%s\"\n", Intersection([2]int{1, 2}, [2]int{2, 3}))
	fmt.Printf("intersection((-1, 1), (0, 4)) ==> \"%s\"\n", Intersection([2]int{-1, 1}, [2]int{0, 4}))
	fmt.Printf("intersection((-3, -1), (-5, 5)) ==> \"%s\"\n", Intersection([2]int{-3, -1}, [2]int{-5, 5}))
}
