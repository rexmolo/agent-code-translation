package main

import (
	"strconv"
)

func ChangeBase(x int, base int) string {
	if x == 0 {
		return ""
	}
	ret := ""
	for x > 0 {
		ret = strconv.Itoa(x%base) + ret
		x /= base
	}
	return ret
}