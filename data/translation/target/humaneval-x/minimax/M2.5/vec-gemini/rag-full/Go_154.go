package main

import (
	"fmt"
	"strings"
)

func CycpatternCheck(a, b string) bool {
	l := len(b)
	if l > len(a) {
		return false
	}
	pat := b + b
	for i := 0; i <= len(a)-l; i++ {
		for j := 0; j <= l; j++ {
			if a[i:i+l] == pat[j:j+l] {
				return true
			}
		}
	}
	return false
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	if CycpatternCheck(a, b) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
