package main

import (
	"strings"
)

func Encrypt(s string) string {
	d := "abcdefghijklmnopqrstuvwxyz"
	out := ""
	for _, c := range s {
		idx := strings.Index(d, string(c))
		if idx != -1 {
			newIdx := (idx + 4) % 26
			out += string(d[newIdx])
		} else {
			out += string(c)
		}
	}
	return out
}