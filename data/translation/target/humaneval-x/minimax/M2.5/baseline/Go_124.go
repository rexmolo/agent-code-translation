package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidDate(date string) bool {
	date = strings.TrimSpace(date)

	if date == "" {
		return false
	}

	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return false
	}

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

	if month < 1 || month > 12 {
		return false
	}

	// Note: The original Python code has operator precedence issues
	// but we preserve the original logic as written
	if (month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12) && (day < 1 || day > 31) {
		return false
	}

	if (month == 4 || month == 6 || month == 9 || month == 11) && (day < 1 || day > 30) {
		return false
	}

	if month == 2 && (day < 1 || day > 29) {
		return false
	}

	return true
}

func main() {
	// Test cases from the examples
	testCases := []struct {
		date     string
		expected bool
	}{
		{"03-11-2000", true},
		{"15-01-2012", false},
		{"04-0-2040", false},
		{"06-04-2020", true},
		{"06/04-2020", false},
	}

	for _, tc := range testCases {
		result := ValidDate(tc.date)
		fmt.Printf("ValidDate('%s') => %v (expected: %v)\n", tc.date, result, tc.expected)
	}
}
