package main

import (
	"strconv"
	"strings"
)

// ValidDate validates a given date string and returns true if the date is valid.
// The date is valid if all of the following rules are satisfied:
// 1. The date string is not empty.
// 2. The number of days is not less than 1 or higher than 31 days for months 1,3,5,7,8,10,12.
//    And the number of days is not less than 1 or higher than 30 days for months 4,6,9,11.
//    And, the number of days is not less than 1 or higher than 29 for the month 2.
// 3. The months should not be less than 1 or higher than 12.
// 4. The date should be in the format: mm-dd-yyyy
func ValidDate(date string) bool {
	date = strings.TrimSpace(date)
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

	_, err = strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	is31DayMonth := month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12
	if is31DayMonth && (day < 1 || day > 31) {
		return false
	}

	is30DayMonth := month == 4 || month == 6 || month == 9 || month == 11
	if is30DayMonth && (day < 1 || day > 30) {
		return false
	}

	if month == 2 && (day < 1 || day > 29) {
		return false
	}

	return true
}
