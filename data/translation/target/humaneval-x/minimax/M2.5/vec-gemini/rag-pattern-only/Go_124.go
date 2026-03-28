package main

import (
	"fmt"
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

	var month, day, year int
	_, err := fmt.Sscanf(parts[0], "%d", &month)
	if err != nil {
		return false
	}
	_, err = fmt.Sscanf(parts[1], "%d", &day)
	if err != nil {
		return false
	}
	_, err = fmt.Sscanf(parts[2], "%d", &year)
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
	}

	return true
}
