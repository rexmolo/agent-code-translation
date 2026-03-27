package main

import (
	"fmt"
	"math"
)

func poly(xs []int, x float64) float64 {
	result := 0.0
	for i, coeff := range xs {
		result += float64(coeff) * math.Pow(x, float64(i))
	}
	return result
}

func FindZero(xs []int) float64 {
	begin, end := -1.0, 1.0
	for poly(xs, begin)*poly(xs, end) > 0 {
		begin *= 2.0
		end *= 2.0
	}
	for end-begin > 1e-10 {
		center := (begin + end) / 2.0
		if poly(xs, center)*poly(xs, begin) > 0 {
			begin = center
		} else {
			end = center
		}
	}
	return begin
}

func main() {
	fmt.Printf("%.2f\n", FindZero([]int{1, 2}))
	fmt.Printf("%.2f\n", FindZero([]int{-6, 11, -6, 1}))
}