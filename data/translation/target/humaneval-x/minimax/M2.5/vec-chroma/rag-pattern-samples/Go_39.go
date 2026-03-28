package main

import (
	"fmt"
	"math"
)

func isPrime(p int) bool {
	if p < 2 {
		return false
	}
	sqrtP := int(math.Sqrt(float64(p)))
	for k := 2; k <= sqrtP && k < p-1; k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

func PrimeFib(n int) int {
	f := []int{0, 1}
	for {
		next := f[len(f)-1] + f[len(f)-2]
		f = append(f, next)
		if isPrime(next) {
			n--
		}
		if n == 0 {
			return next
		}
	}
}

func main() {
	// Test cases
	fmt.Println(PrimeFib(1)) // 2
	fmt.Println(PrimeFib(2)) // 3
	fmt.Println(PrimeFib(3)) // 5
	fmt.Println(PrimeFib(4)) // 13
	fmt.Println(PrimeFib(5)) // 89
}
