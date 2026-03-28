package main

import (
	"strconv"
	"strings"
)

func ValidDate(date string) bool {
	// Trim whitespace
	date = strings.TrimSpace(date)
	
	// Check if empty
	if date == "" {
		return false
	}
	
	// Check format mm-dd-yyyy (must have exactly 2 dashes)
	parts := strings.Split(date, "-")
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
	
	// Validate month range (1-12)
	if month < 1 || month > 12 {
		return false
	}
	
	// Validate days based on month - using proper parentheses for operator precedence
	// Months with 31 days: 1, 3, 5, 7, 8, 10, 12
	if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) && (day < 1 || day > 31) {
		return false
	}
	// Months with 30 days: 4, 6, 9, 11
	if (month == 4 || month == 6 || month == 9 || month == 11) && (day < 1 || day > 30) {
		return false
	}
	// Month 2 (February) has 29 days max
	if month == 2 && (day < 1 || day > 29) {
		return false
	}
	
	return true
}
