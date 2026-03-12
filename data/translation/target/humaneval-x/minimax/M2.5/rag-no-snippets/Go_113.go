package main

import (
	"fmt"
	"strconv"
)

func OddCount(lst []string) []string {
	res := make([]string, 0, len(lst))
	for _, arr := range lst {
		n := 0
		for _, d := range arr {
			digit := int(d - '0')
			if digit%2 == 1 {
				n++
			}
		}
		res = append(res, fmt.Sprintf("the number of odd elements %dn the str%dnng %d of the %dnput.", n, n, n, n))
	}
	return res
}