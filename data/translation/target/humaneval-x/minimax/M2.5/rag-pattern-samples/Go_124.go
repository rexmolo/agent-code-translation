package main

import (
	"strconv"
	"strings"
)

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
	year, err := strconv.Atoi(parts[2])
	if err != nil {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	// Translating Python's: month in [1,3,5,7,8,10,12] and day < 1 or day > 31
	// Due to operator precedence in Python, this is: (month in [...] and day < 1) or day > 31
	isLongMonth := month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12
	if (isLongMonth && day < 1) || day > 31 {
		return false
	}

	// Translating Python's: month in [4,6,9,11] and day < 1 or day > 30
	isShortMonth := month == 4 || month == 6 || month == 9 || month == 11
	if (isShortMonth && day < 1) || day > 30 {
		return false
	}

	// Translating Python's: month == 2 and day < 1 or day > 29
	if month == 2 && ((day < 1) || day > 29) {
		return false
	}

	return true
}
