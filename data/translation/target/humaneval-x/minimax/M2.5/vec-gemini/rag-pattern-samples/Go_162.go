package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func StringToMd5(text string) interface{} {
	if text == "" {
		return nil
	}
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	// Test cases
	testCases := []string{"Hello world", "", "test"}
	for _, tc := range testCases {
		result := StringToMd5(tc)
		if result == nil {
			fmt.Println("Input: '\"' -> Output: nil")
		} else {
			fmt.Printf("Input: '%s' -> Output: '%s'\n", tc, result)
		}
	}
}
