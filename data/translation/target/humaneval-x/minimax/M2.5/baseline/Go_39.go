package main

import (
	"fmt"
	"math"
)

func isPrime(p int) bool {
	if p < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(p))) + 1
	for k := 2; k < limit && k < p; k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

func PrimeFib(n int) int {
	f := []int{0, 1}
	for {
		// Append the next Fibonacci number
		f = append(f, f[len(f)-1]+f[len(f)-2])
		if isPrime(f[len(f)-1]) {
			n--
		}
		if n == 0 {
			return f[len(f)-1]
		}
	}
}

func main() {
	fmt.Println(PrimeFib(1)) // 2
	fmt.Println(PrimeFib(2)) // 3
	fmt.Println(PrimeFib(3)) // 5
	fmt.Println(PrimeFib(4)) // 13
	fmt.Println(PrimeFib(5)) // 89
}