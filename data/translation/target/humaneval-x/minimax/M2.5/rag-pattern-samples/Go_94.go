package main

import "fmt"

func Skjkasdkd(lst []int) int {
	// Helper function to check if a number is prime
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
	for _, v := range lst {
		if v > maxx && isPrime(v) {
			maxx = v
		}
	}

	// Calculate sum of digits
	result := 0
	temp := maxx
	if temp == 0 {
		return 0
	}
	for temp > 0 {
		result += temp % 10
		temp /= 10
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(Skjkasdkd([]int{0, 3, 2, 1, 3, 5, 7, 4, 5, 5, 5, 2, 181, 32, 4, 32, 3, 2, 32, 324, 4, 3}))
	fmt.Println(Skjkasdkd([]int{1, 0, 1, 8, 2, 4597, 2, 1, 3, 40, 1, 2, 1, 2, 4, 2, 5, 1}))
	fmt.Println(Skjkasdkd([]int{1, 3, 1, 32, 5107, 34, 83278, 109, 163, 23, 2323, 32, 30, 1, 9, 3}))
	fmt.Println(Skjkasdkd([]int{0, 724, 32, 71, 99, 32, 6, 0, 5, 91, 83, 0, 5, 6}))
	fmt.Println(Skjkasdkd([]int{0, 81, 12, 3, 1, 21}))
	fmt.Println(Skjkasdkd([]int{0, 8, 1, 2, 1, 7}))
}
