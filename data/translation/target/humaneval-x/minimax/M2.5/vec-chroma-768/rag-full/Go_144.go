package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split both fractions
	partsX := strings.Split(x, "/")
	partsN := strings.Split(n, "/")

	// Convert to integers
	a, err := strconv.Atoi(partsX[0])
	if err != nil {
		return false
	}
	b, err := strconv.Atoi(partsX[1])
	if err != nil {
		return false
	}
	c, err := strconv.Atoi(partsN[0])
	if err != nil {
		return false
	}
	d, err := strconv.Atoi(partsN[1])
	if err != nil {
		return false
	}

	// Multiply numerators and denominators
	numerator := a * c
	denom := b * d

	// Check if result is a whole number (no remainder)
	return numerator%denom == 0
}

func main() {
	// Test examples
	fmt.Println(Simplify("1/5", "5/1")) // true
	fmt.Println(Simplify("1/6", "2/1")) // false
	fmt.Println(Simplify("7/10", "10/2")) // false
}