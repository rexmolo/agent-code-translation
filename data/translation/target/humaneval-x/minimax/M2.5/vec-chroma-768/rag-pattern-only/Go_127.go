package main

import "fmt"

func Intersection(interval1 [2]int, interval2 [2]int) string {
	// Helper function to check if a number is prime
	isPrime := func(num int) bool {
		if num <= 1 {
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
	l := interval1[0]
	if interval2[0] > l {
		l = interval2[0]
	}
	r := interval1[1]
	if interval2[1] < r {
		r = interval2[1]
	}

	length := r - l

	if length > 0 && isPrime(length) {
		return "YES"
	}
	return "NO"
}

func main() {
	// Test the function with sample cases
	fmt.Println(Intersection([2]int{1, 2}, [2]int{2, 3}))        // NO
	fmt.Println(Intersection([2]int{-1, 1}, [2]int{0, 4}))       // NO
	fmt.Println(Intersection([2]int{-3, -1}, [2]int{-5, 5}))     // YES
}
