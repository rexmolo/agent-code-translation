package main

import (
	"math"
	"strconv"
)

func Skjkasdkd(lst []int) int {
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
	for _, v := range lst {
		if v > maxx && isPrime(v) {
			maxx = v
		}
	}

	result := 0
	for _, c := range strconv.Itoa(maxx) {
		result += int(c - '0')
	}

	return result
}

func main() {
	// Example test cases
	lst1 := []int{0, 3, 2, 1, 3, 5, 7, 4, 5, 5, 5, 2, 181, 32, 4, 32, 3, 2, 32, 324, 4, 3}
	lst2 := []int{1, 0, 1, 8, 2, 4597, 2, 1, 3, 40, 1, 2, 1, 2, 4, 2, 5, 1}
	lst3 := []int{1, 3, 1, 32, 5107, 34, 83278, 109, 163, 23, 2323, 32, 30, 1, 9, 3}
	lst4 := []int{0, 724, 32, 71, 99, 32, 6, 0, 5, 91, 83, 0, 5, 6}
	lst5 := []int{0, 81, 12, 3, 1, 21}
	lst6 := []int{0, 8, 1, 2, 1, 7}

	println(Skjkasdkd(lst1)) // 10
	println(Skjkasdkd(lst2)) // 25
	println(Skjkasdkd(lst3)) // 13
	println(Skjkasdkd(lst4)) // 11
	println(Skjkasdkd(lst5)) // 3
	println(Skjkasdkd(lst6)) // 7
}