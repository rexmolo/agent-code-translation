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
	fmt.Println(Encrypt("hi"))
	fmt.Println(Encrypt("asdfghjkl"))
	fmt.Println(Encrypt("gf"))
	fmt.Println(Encrypt("et"))
}
