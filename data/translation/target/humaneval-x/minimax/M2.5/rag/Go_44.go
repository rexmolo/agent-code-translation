package main

import (
	"strconv"
)

func ChangeBase(x int, base int) string {
	ret := ""
	for x > 0 {
		ret = strconv.Itoa(x%base) + ret
		x /= base
	}
	return ret
}