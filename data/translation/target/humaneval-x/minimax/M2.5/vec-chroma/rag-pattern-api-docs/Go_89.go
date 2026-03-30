package main

import (
	"fmt"
	"slices"
	"strings"
)

func Encrypt(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder

	for _, c := range s {
		idx := slices.Index([]byte(alphabet), byte(c))
		if idx != -1 {
			result.WriteByte(alphabet[(idx+4)%26])
		} else {
			result.WriteByte(byte(c))
		}
	}

	return result.String()
}

func main() {
	fmt.Println(Encrypt("hi"))         // lm
	fmt.Println(Encrypt("asdfghjkl")) // ewhjklnop
	fmt.Println(Encrypt("gf"))         // kj
	fmt.Println(Encrypt("et"))          // ix
}
