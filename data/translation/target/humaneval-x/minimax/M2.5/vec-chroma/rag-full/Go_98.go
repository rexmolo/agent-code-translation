package main

import "fmt"

func CountUpper(s string) int {
	count := 0
	vowels := map[rune]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for i := 0; i < len(s); i += 2 {
		if vowels[rune(s[i])] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountUpper("aBCdEf"))
	fmt.Println(CountUpper("abcdefg"))
	fmt.Println(CountUpper("dBBE"))
}
