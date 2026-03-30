package main

import (
	"strings"
	"strconv"
)

func Simplify(x, n string) bool {
	// Split the first fraction
	parts1 := strings.Split(x, "/")
	a := parts1[0]
	b := parts1[1]

	// Split the second fraction
	parts2 := strings.Split(n, "/")
	c := parts2[0]
	d := parts2[1]

	// Convert strings to integers
	aNum, err := strconv.Atoi(a)
	if err != nil {
		return false
	}
	bNum, err := strconv.Atoi(b)
	if err != nil {
		return false
	}
	cNum, err := strconv.Atoi(c)
	if err != nil {
		return false
	}
	dNum, err := strconv.Atoi(d)
	if err != nil {
		return false
	}

	// Calculate the product of fractions
	numerator := aNum * cNum
	denom := bNum * dNum

	// Check if the result is a whole number (no remainder)
	return numerator%denom == 0
}
