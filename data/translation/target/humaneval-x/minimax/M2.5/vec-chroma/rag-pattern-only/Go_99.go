package main

import (
	"math"
	"strconv"
	"strings"
)

func ClosestInteger(value string) int {
	if strings.Contains(value, ".") {
		// Remove trailing zeros
		value = strings.TrimRight(value, "0")
		// Also remove trailing dot if present after removing zeros
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	} else if len(value) > 0 {
		res := int(math.Round(num))
		return res
	}

	return 0
}