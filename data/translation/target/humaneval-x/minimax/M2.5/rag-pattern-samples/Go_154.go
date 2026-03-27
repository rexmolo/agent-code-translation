package main

import "fmt"

func CycpatternCheck(a, b string) bool {
	l := len(b)
	if l == 0 {
		return true
	}
	pat := b + b

	for i := 0; i <= len(a)-l; i++ {
		for j := 0; j < l; j++ {
			if a[i:i+l] == pat[j:j+l] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(CycpatternCheck("abcd", "abd"))   // False
	fmt.Println(CycpatternCheck("hello", "ell")) // True
	fmt.Println(CycpatternCheck("whassup", "psus")) // False
	fmt.Println(CycpatternCheck("abab", "baa"))  // True
	fmt.Println(CycpatternCheck("efef", "eeff")) // False
	fmt.Println(CycpatternCheck("himenss", "simen")) // True
}