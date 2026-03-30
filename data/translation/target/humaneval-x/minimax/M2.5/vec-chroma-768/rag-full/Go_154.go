package main

import "fmt"

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
	fmt.Println(CycpatternCheck("abcd", "abd"))   // false
	fmt.Println(CycpatternCheck("hello", "ell"))  // true
	fmt.Println(CycpatternCheck("whassup", "psus")) // false
	fmt.Println(CycpatternCheck("abab", "baa"))   // true
	fmt.Println(CycpatternCheck("efef", "eeff"))  // false
	fmt.Println(CycpatternCheck("himenss", "simen")) // true
}
