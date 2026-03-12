package main

import (
	"fmt"
	"math"
)

func Skjkasdkd(lst []int) int {
	// isPrime function - checks if n is a prime number
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	maxx := 0
	for _, val := range lst {
		if val > maxx && isPrime(val) {
			maxx = val
		}
	}

	// Sum the digits of maxx
	result := 0
	temp := maxx
	for temp > 0 {
		result += temp % 10
		temp /= 10
	}

	return result
}

func main() {
	// Test cases
	lst1 := []int{0, 3, 2, 1, 3, 5, 7, 4, 5, 5, 5, 2, 181, 32, 4, 32, 3, 2, 32, 324, 4, 3}
	fmt.Println(Skjkasdkd(lst1)) // Expected: 10

	lst2 := []int{1, 0, 1, 8, 2, 4597, 2, 1, 3, 40, 1, 2, 1, 2, 4, 2, 5, 1}
	fmt.Println(Skjkasdkd(lst2)) // Expected: 25

	lst3 := []int{1, 3, 1, 32, 5107, 34, 83278, 109, 163, 23, 2323, 32, 30, 1, 9, 3}
	fmt.Println(Skjkasdkd(lst3)) // Expected: 13

	lst4 := []int{0, 724, 32, 71, 99, 32, 6, 0, 5, 91, 83, 0, 5, 6}
	fmt.Println(Skjkasdkd(lst4)) // Expected: 11

	lst5 := []int{0, 81, 12, 3, 1, 21}
	fmt.Println(Skjkasdkd(lst5)) // Expected: 3

	lst6 := []int{0, 8, 1, 2, 1, 7}
	fmt.Println(Skjkasdkd(lst6)) // Expected: 7
}
