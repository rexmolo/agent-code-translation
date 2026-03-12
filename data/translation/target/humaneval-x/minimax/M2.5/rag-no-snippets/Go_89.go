package main

import (
	"fmt"
	"strings"
)

func Encrypt(s string) string {
	d := "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder
	for _, c := range s {
		idx := strings.Index(d, string(c))
		if idx != -1 {
			newIdx := (idx + 4) % 26
			result.WriteByte(d[newIdx])
		} else {
			result.WriteByte(c)
		}
	}
	return result.String()
}

func main() {
	fmt.Println(Encrypt("hi"))          // lm
	fmt.Println(Encrypt("asdfghjkl"))   // ewhjklnop
	fmt.Println(Encrypt("gf"))          // kj
	fmt.Println(Encrypt("et"))          // ix
}
