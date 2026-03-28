package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split both fractions into numerator and denominator
	partsX := strings.Split(x, "/")
	partsN := strings.Split(n, "/")

	// Parse all values as integers
	a, _ := strconv.Atoi(partsX[0])
	b, _ := strconv.Atoi(partsX[1])
	c, _ := strconv.Atoi(partsN[0])
	d, _ := strconv.Atoi(partsN[1])

	// Compute the product: (a/b) * (c/d) = (a*c) / (b*d)
	numerator := a * c
	denom := b * d

	// Check if the result is a whole number (no remainder)
	if numerator%denom == 0 {
		return true
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(Simplify("1/5", "5/1"))   // True
	fmt.Println(Simplify("1/6", "2/1"))   // False
	fmt.Println(Simplify("7/10", "10/2")) // False
}