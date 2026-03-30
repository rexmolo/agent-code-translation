package main

import (
	"strconv"
	"strings"
)

// ValidDate validates a given date string and returns true if the date is valid,
// otherwise false. The date is valid if all of the following rules are satisfied:
// 1. The date string is not empty.
// 2. The number of days is valid for the given month.
// 3. The months should be between 1 and 12.
// 4. The date should be in the format: mm-dd-yyyy
func ValidDate(date string) bool {
	// Strip whitespace
	date = strings.TrimSpace(date)

	// Check if empty
	if date == "" {
		return false
	}

	// Split by '-'
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return false
	}

	// Convert to integers
	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}
	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}
	_, err = strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	// Validate month (1-12)
	if month < 1 || month > 12 {
		return false
	}

	// Define months with 31 days
	monthsWith31Days := map[int]bool{
		1: true, 3: true, 5: true, 7: true, 8: true, 10: true, 12: true,
	}

	// Define months with 30 days
	monthsWith30Days := map[int]bool{
		4: true, 6: true, 9: true, 11: true,
	}

	// Validate day based on month
	if monthsWith31Days[month] {
		if day < 1 || day > 31 {
			return false
		}
	} else if monthsWith30Days[month] {
		if day < 1 || day > 30 {
			return false
		}
	} else if month == 2 { // February
		if day < 1 || day > 29 {
			return false
		}
	}

	return true
}
