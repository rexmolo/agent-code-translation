package main

import (
	"strconv"
	"strings"
)

func ValidDate(date string) bool {
	// Strip whitespace from the date string
	date = strings.TrimSpace(date)

	// Check if the date string is empty
	if date == "" {
		return false
	}

	// Split the date by '-'
	parts := strings.Split(date, "-")

	// Check if we have exactly 3 parts (mm-dd-yyyy format)
	if len(parts) != 3 {
		return false
	}

	// Convert each part to integer
	month, err := strconv.Atoi(parts[0])
	if err != nil {
		return false
	}

	day, err := strconv.Atoi(parts[1])
	if err != nil {
		return false
	}

	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	// Check if month is in valid range (1-12)
	if month < 1 || month > 12 {
		return false
	}

	// Define months with 31 days: 1,3,5,7,8,10,12
	isMonthWith31Days := func(m int) bool {
		switch m {
		case 1, 3, 5, 7, 8, 10, 12:
			return true
		}
		return false
	}

	// Check day validity based on month
	if isMonthWith31Days(month) {
		if day < 1 || day > 31 {
			return false
		}
	} else if month == 2 {
		// February can have 1-29 days
		if day < 1 || day > 29 {
			return false
		}
	} else {
		// Months with 30 days: 4,6,9,11
		if day < 1 || day > 30 {
			return false
		}
	}

	return true
}