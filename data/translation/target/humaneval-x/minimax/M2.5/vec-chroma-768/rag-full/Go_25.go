package main

import (
	"fmt"
	"math"
)

func Factorize(n int) []int {
	fact := make([]int, 0)
	for i := 2; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			fact = append(fact, i)
			n /= i
			i--
		}
	}

	if n > 1 {
		fact = append(fact, n)
	}
	return fact
}

func main() {
	fmt.Println(Factorize(8))
	fmt.Println(Factorize(25))
	fmt.Println(Factorize(70))
}
