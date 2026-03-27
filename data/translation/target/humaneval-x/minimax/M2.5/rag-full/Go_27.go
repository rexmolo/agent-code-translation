package main

import "fmt"

// FlipCase flips the case of each character in the string:
// lowercase characters become uppercase and uppercase become lowercase.
func FlipCase(str string) string {
	result := make([]rune, len(str))
	for i, c := range str {
		if c >= 'a' && c <= 'z' {
			result[i] = c - 'a' + 'A'
		} else if c >= 'A' && c <= 'Z' {
			result[i] = c - 'A' + 'a'
		} else {
			result[i] = c
		}
	}
	return string(result)
}

func main() {
	fmt.Println(FlipCase("Hello")) // Output: hELLO
}
