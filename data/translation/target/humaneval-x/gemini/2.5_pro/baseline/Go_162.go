package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// StringToMd5 calculates the MD5 hash of a string.
// It returns the hex-encoded hash as a string.
// If the input text is an empty string, it returns nil.
func StringToMd5(text string) interface{} {
	if text == "" {
		return nil
	}
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	// Demonstrate the function with the example from the Python docstring.
	hash1 := StringToMd5("Hello world")
	fmt.Printf("'Hello world' -> %v\n", hash1)

	// Demonstrate the empty string case.
	hash2 := StringToMd5("")
	fmt.Printf("'' -> %v\n", hash2)

	// Verification
	expectedHash := "3e25960a79dbc69b674cd4ec67a72c62"
	if hash1 == expectedHash {
		fmt.Println("Test case 'Hello world' passed.")
	} else {
		fmt.Println("Test case 'Hello world' failed.")
	}

	if hash2 == nil {
		fmt.Println("Test case for empty string passed.")
	} else {
		fmt.Println("Test case for empty string failed.")
	}
}
