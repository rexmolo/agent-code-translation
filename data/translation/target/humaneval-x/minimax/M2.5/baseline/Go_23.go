package main

import "fmt"

func Strlen(str string) int {
	return len(str)
}

func main() {
	// Test cases
	fmt.Println(Strlen(""))    // 0
	fmt.Println(Strlen("abc")) // 3
}
