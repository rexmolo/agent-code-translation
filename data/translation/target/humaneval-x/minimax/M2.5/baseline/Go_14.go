package main

import (
	"fmt"
)

func AllPrefixes(str string) []string {
	result := make([]string, 0)

	for i := 0; i < len(str); i++ {
		result = append(result, str[:i+1])
	}

	return result
}

func main() {
	// Test the function
	fmt.Println(AllPrefixes("abc"))
}
