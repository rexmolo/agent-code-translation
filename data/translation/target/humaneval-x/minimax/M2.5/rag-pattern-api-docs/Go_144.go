package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split the first fraction x
	partsX := strings.Split(x, "/")
	a, err := strconv.Atoi(partsX[0])
	if err != nil {
		return false
	}
	b, err := strconv.Atoi(partsX[1])
	if err != nil {
		return false
	}

	// Split the second fraction n
	partsN := strings.Split(n, "/")
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

	// Check if x * n evaluates to a whole number
	// A whole number means numerator is divisible by denominator
	return numerator%denom == 0
}

func main() {
	// Test cases
	fmt.Println(Simplify("1/5", "5/1"))  // true
	fmt.Println(Simplify("1/6", "2/1"))  // false
	fmt.Println(Simplify("7/10", "10/2")) // false
}
