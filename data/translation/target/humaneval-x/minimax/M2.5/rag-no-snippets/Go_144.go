package main

import (
	"strconv"
	"strings"
)

func Simplify(x, n string) bool {
	// Split x and n by "/"
	xParts := strings.Split(x, "/")
	nParts := strings.Split(n, "/")

	// Parse numerators and denominators
	a, _ := strconv.ParseInt(xParts[0], 10, 64)
	b, _ := strconv.ParseInt(xParts[1], 10, 64)
	c, _ := strconv.ParseInt(nParts[0], 10, 64)
	d, _ := strconv.ParseInt(nParts[1], 10, 64)

	// Multiply numerators and denominators
	numerator := a * c
	denom := b * d

	// Check if the result is a whole number (divisible evenly)
	if numerator%denom == 0 {
		return true
	}
	return false
}