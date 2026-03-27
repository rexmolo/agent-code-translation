package main

import "fmt"

func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		switch s[i] {
		case 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountUpper("aBCdEf")) // 1
	fmt.Println(CountUpper("abcdefg")) // 0
	fmt.Println(CountUpper("dBBE")) // 0
}
