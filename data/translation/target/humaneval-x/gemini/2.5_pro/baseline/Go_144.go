package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Simplify determines if the product of two fractions, x and n, results in a whole number.
// Both x and n are given as string representations of fractions, e.g., "numerator/denominator".
func Simplify(x, n string) bool {
	// Split the fraction strings into numerator and denominator parts.
	partsX := strings.Split(x, "/")
	partsN := strings.Split(n, "/")

	// Convert string parts to integers. The problem guarantees valid input,
	// so we can use the blank identifier `_` to ignore potential errors from Atoi.
	a, _ := strconv.Atoi(partsX[0])
	b, _ := strconv.Atoi(partsX[1])
	c, _ := strconv.Atoi(partsN[0])
	d, _ := strconv.Atoi(partsN[1])

	// Calculate the numerator and denominator of the resulting fraction (a/b * c/d = ac/bd).
	numerator := a * c
	denominator := b * d

	// A fraction evaluates to a whole number if its numerator is perfectly
	// divisible by its denominator. The modulo operator checks for this.
	return numerator%denominator == 0
}

func main() {
	// Example cases from the problem description.
	fmt.Println(Simplify("1/5", "5/1"))
	fmt.Println(Simplify("1/6", "2/1"))
	fmt.Println(Simplify("7/10", "10/2"))
	fmt.Println(Simplify("2/4", "2/1")) // 4/4 = 1 -> True
}
