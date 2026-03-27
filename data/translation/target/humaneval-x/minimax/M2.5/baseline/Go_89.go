package main

import (
	"strings"
)

func Encrypt(s string) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder
	
	for _, c := range s {
		idx := strings.Index(alphabet, string(c))
		if idx != -1 {
			newIdx := (idx + 4) % 26
			result.WriteByte(alphabet[newIdx])
		} else {
			result.WriteRune(c)
		}
	}
	
	return result.String()
}
