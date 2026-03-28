package main

import (
	"math"
	"strconv"
)

func Skjkasdkd(lst []int) int {
	// Helper function to check if a number is prime
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	maxx := 0
	for i := 0; i < len(lst); i++ {
		if lst[i] > maxx && isPrime(lst[i]) {
			maxx = lst[i]
		}
	}

	// Sum the digits of maxx
	result := 0
	for _, c := range strconv.Itoa(maxx) {
		result += int(c - '0')
	}

	return result
}