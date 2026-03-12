package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		for strings.HasSuffix(value, "0") {
			value = value[:len(value)-1]
		}
	}

	// Handle case where value becomes "." or ends with "."
	if strings.HasSuffix(value, ".") {
		value = value[:len(value)-1]
	}

	num, _ := strconv.ParseFloat(value, 64)

	// Check if the cleaned string ends with ".5"
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			res := math.Ceil(num)
			return int(res)
		} else {
			res := math.Floor(num)
			return int(res)
		}
	} else if len(value) > 0 {
		res := math.Round(num)
		return int(res)
	}

	return 0
}
