package main

import "strconv"

func ChangeBase(x int, base int) string {
	result := ""
	for x > 0 {
		result = strconv.Itoa(x%base) + result
		x /= base
	}
	return result
}