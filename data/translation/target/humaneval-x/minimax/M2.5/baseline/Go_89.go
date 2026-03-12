package main

import (
	"fmt"
	"strings"
)

func Encrypt(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	shift := 2 * 2 // 4 places

	var result strings.Builder
	for _, c := range s {
		idx := strings.Index(alphabet, string(c))
		if idx != -1 {
			newIdx := (idx + shift) % 26
			result.WriteByte(alphabet[newIdx])
		} else {
			result.WriteRune(c)
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