package main

import (
    "math"
)

// poly evaluates polynomial with coefficients xs at point x.
// Returns xs[0] + xs[1] * x + xs[2] * x^2 + .... xs[n] * x^n
func poly(xs []int, x float64) float64 {
    result := 0.0
    for i, coeff := range xs {
        result += float64(coeff) * math.Pow(x, float64(i))
    }
    return result
}

// FindZero finds x such that poly(x) = 0.
// FindZero returns only one zero point, even if there are many.
// Moreover, FindZero only takes a slice xs having even number of coefficients
// and largest non zero coefficient as it guarantees a solution.
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
    // Test cases from Python docstrings
    result1 := FindZero([]int{1, 2})
    result2 := FindZero([]int{-6, 11, -6, 1})
    println("find_zero([1, 2]):", result1)
    println("find_zero([-6, 11, -6, 1]):", result2)
}
