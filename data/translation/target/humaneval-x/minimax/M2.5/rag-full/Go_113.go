package main

import (
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
		res = append(res, "the number of odd elements "+strconv.Itoa(n)+"n the str"+strconv.Itoa(n)+"ng "+strconv.Itoa(n)+" of the "+strconv.Itoa(n)+"nput.")
	}
	return res
}