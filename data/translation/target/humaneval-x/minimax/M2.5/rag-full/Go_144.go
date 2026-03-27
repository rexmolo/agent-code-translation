package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split the first fraction x
	parts := strings.Split(x, "/")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])

	// Split the second fraction n
	parts = strings.Split(n, "/")
	c, _ := strconv.Atoi(parts[0])
	d, _ := strconv.Atoi(parts[1])

	// Calculate numerator and denominator of the product
	numerator := a * c
	denom := b * d

	// Check if the result is a whole number
	// Using modulo to check if numerator divides evenly by denominator
	return numerator%denom == 0
}

func main() {
	// Test cases
	fmt.Println(Simplify("1/5", "5/1"))   // Expected: true
	fmt.Println(Simplify("1/6", "2/1"))   // Expected: false
	fmt.Println(Simplify("7/10", "10/2")) // Expected: false
}
