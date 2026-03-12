package main

import (
	"strconv"
	"strings"
)

// ValidDate validates a given date string according to specific rules.
// The date is valid if all of the following rules are satisfied:
// 1. The date string is not empty.
// 2. The day/month ranges are valid (without leap year considerations).
// 3. The months should be between 1 and 12.
// 4. The date should be in the format: mm-dd-yyyy
func ValidDate(date string) bool {
	date = strings.TrimSpace(date)
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return false
	}

	monthStr, dayStr, yearStr := parts[0], parts[1], parts[2]

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return false
	}

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return false
	}

	// The year is parsed only to ensure it's a number, matching the Python logic.
	_, err = strconv.Atoi(yearStr)
	if err != nil {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day < 1 || day > 31 {
			return false
		}
	case 4, 6, 9, 11:
		if day < 1 || day > 30 {
			return false
		}
	case 2:
		if day < 1 || day > 29 {
			return false
		}
	default:
		// This case is technically unreachable due to the month check above.
		return false
	}

	return true
}
