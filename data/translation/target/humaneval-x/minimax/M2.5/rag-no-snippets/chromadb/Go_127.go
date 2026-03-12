package main

import (
	"slices"
)

func isPrime(num int) bool {
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

func Intersection(interval1 [2]int, interval2 [2]int) string {
	l := slices.Max([]int{interval1[0], interval2[0]})
	r := slices.Min([]int{interval1[1], interval2[1]})
	length := r - l
	if length > 0 && isPrime(length) {
		return "YES"
	}
	return "NO"
}

func main() {
	// Test cases
	// fmt.Println(Intersection([2]int{1, 2}, [2]int{2, 3}))      // NO
	// fmt.Println(Intersection([2]int{-1, 1}, [2]int{0, 4}))    // NO
	// fmt.Println(Intersection([2]int{-3, -1}, [2]int{-5, 5}))  // YES
}
