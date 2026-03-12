package main

import (
	"math"
	"strconv"
)

func ClosestInteger(value string) int {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	// Go's math.Round function rounds to the nearest integer, with .5 cases
	// rounded away from zero, which is the exact behavior required.
	// For example, Round(14.5) is 15 and Round(-14.5) is -15.
	rounded := math.Round(num)

	return int(rounded)
}
