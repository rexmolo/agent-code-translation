package main

import (
	"math"
	"strings"
)

func ClosestInteger(value string) int {
	// Remove trailing zeros after decimal point
	if strings.Contains(value, ".") {
		value = strings.TrimRight(value, "0")
		// If it ends with ".", remove it
		if strings.HasSuffix(value, ".") {
			value = value[:len(value)-1]
		}
	}

	// Handle empty string case
	if len(value) == 0 {
		return 0
	}

	num := parseFloat(value)

	// Check if value ends with ".5"
	if strings.HasSuffix(value, ".5") {
		if num > 0 {
			return int(math.Ceil(num))
		} else {
			return int(math.Floor(num))
		}
	}

	// Use round-half-away-from-zero
	return roundHalfAwayFromZero(num)
}

func parseFloat(s string) float64 {
	// Simple float parser for basic cases
	var negative bool
	var hasDecimal bool
	var integerPart string
	var fractionPart string

	if strings.HasPrefix(s, "-") {
		negative = true
		s = s[1:]
	} else if strings.HasPrefix(s, "+") {
		s = s[1:]
	}

	parts := strings.Split(s, ".")
	integerPart = parts[0]
	if len(parts) > 1 {
		hasDecimal = true
		fractionPart = parts[1]
	}

	var intVal float64 = 0
	for _, c := range integerPart {
		intVal = intVal*10 + float64(c-'0')
	}

	if hasDecimal && len(fractionPart) > 0 {
		fractionVal := 0.0
		divisor := 1.0
		for _, c := range fractionPart {
			fractionVal = fractionVal*10 + float64(c-'0')
			divisor *= 10
		}
		intVal += fractionVal / divisor
	}

	if negative {
		return -intVal
	}
	return intVal
}

func roundHalfAwayFromZero(x float64) int {
	if x >= 0 {
		return int(math.Floor(x + 0.5))
	}
	return int(math.Ceil(x - 0.5))
}