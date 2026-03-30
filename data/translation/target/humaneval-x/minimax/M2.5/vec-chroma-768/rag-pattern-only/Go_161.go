package main

import "fmt"

func Solve(s string) string {
	runes := []rune(s)
	flg := 0

	for idx, char := range runes {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if char >= 'a' && char <= 'z' {
				runes[idx] = char - 32 // convert to uppercase (ASCII difference)
			} else {
				runes[idx] = char + 32 // convert to lowercase (ASCII difference)
			}
			flg = 1
		}
	}

	if flg == 0 {
		// Reverse the string
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}

func main() {
	// Test examples
	fmt.Println(Solve("1234"))   // Expected: "4321"
	fmt.Println(Solve("ab"))     // Expected: "AB"
	fmt.Println(Solve("#a@C"))  // Expected: "#A@c"
}
