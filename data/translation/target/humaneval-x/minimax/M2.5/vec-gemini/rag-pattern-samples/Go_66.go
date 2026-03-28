package main

import "fmt"

func Digitsum(x string) int {
	sum := 0
	for _, ch := range x {
		if ch >= 'A' && ch <= 'Z' {
			sum += int(ch)
		}
	}
	return sum
}

func main() {
	// Test cases
	fmt.Println(Digitsum(""))        // 0
	fmt.Println(Digitsum("abAB"))    // 131
	fmt.Println(Digitsum("abcCd"))   // 67
	fmt.Println(Digitsum("helloE"))  // 69
	fmt.Println(Digitsum("woArBld")) // 131
	fmt.Println(Digitsum("aAaaaXa")) // 153
}