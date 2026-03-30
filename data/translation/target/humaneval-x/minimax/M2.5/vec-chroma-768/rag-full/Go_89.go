package main

import (
	"fmt"
	"strings"
)

func Encrypt(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	result := ""
	for _, c := range s {
		idx := strings.Index(alphabet, string(c))
		if idx != -1 {
			result += string(alphabet[(idx+4)%26])
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	fmt.Println(Encrypt("hi"))        // lm
	fmt.Println(Encrypt("asdfghjkl")) // ewhjklnop
	fmt.Println(Encrypt("gf"))        // kj
	fmt.Println(Encrypt("et"))        // ix
}
