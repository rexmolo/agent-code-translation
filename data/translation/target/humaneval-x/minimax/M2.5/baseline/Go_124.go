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
