package main

import (
	"fmt"
	"math"
)

func PrimeFib(n int) int {
	isPrime := func(p int) bool {
		if p < 2 {
			return false
		}
		for k := 2; k <= int(math.Sqrt(float64(p))); k++ {
			if p%k == 0 {
				return false
			}
		}
		return true
	}

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
	fmt.Println(PrimeFib(1))
	fmt.Println(PrimeFib(2))
	fmt.Println(PrimeFib(3))
	fmt.Println(PrimeFib(4))
	fmt.Println(PrimeFib(5))
}