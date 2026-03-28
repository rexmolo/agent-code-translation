package main

import "slices"

func Intersection(interval1 [2]int, interval2 [2]int) string {
	// Helper function to check if a number is prime
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

	// Find the intersection bounds
	l := slices.Max([]int{interval1[0], interval2[0]})
	r := slices.Min([]int{interval1[1], interval2[1]})
	length := r - l

	// Return YES if intersection exists and length is prime
	if length > 0 && isPrime(length) {
		return "YES"
	}
	return "NO"
}
