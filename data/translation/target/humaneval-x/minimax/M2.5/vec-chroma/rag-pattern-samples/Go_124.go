package main

import (
	"strconv"
	"strings"
)

func ValidDate(date string) bool {
	// Strip whitespace
	date = strings.TrimSpace(date)

	// Check if date string is empty
	if date == "" {
		return false
	}

	// Split by '-' - must have exactly 3 parts
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return false
	}

	// Convert to integers
	month, err1 := strconv.Atoi(parts[0])
	day, err2 := strconv.Atoi(parts[1])
	year, err3 := strconv.Atoi(parts[2])

	// If any conversion fails, return false
	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}

	// Validate month range (1-12)
	if month < 1 || month > 12 {
		return false
	}

	// Validate days based on month
	// 31-day months: 1,3,5,7,8,10,12
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		if day < 1 || day > 31 {
			return false
		}
	}
	// 30-day months: 4,6,9,11
	if month == 4 || month == 6 || month == 9 || month == 11 {
		if day < 1 || day > 30 {
			return false
		}
	}
	// February (month 2)
	if month == 2 {
		if day < 1 || day > 29 {
			return false
		}
	}

	return true
}