package main

import (
	"fmt"
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	strong := extensions[0]
	myVal := calculateStrength(extensions[0])

	for _, s := range extensions {
		val := calculateStrength(s)
		if val > myVal {
			strong = s
			myVal = val
		}
	}

	return className + "." + strong
}

func calculateStrength(ext string) int {
	upperCount := 0
	lowerCount := 0
	for _, r := range ext {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				upperCount++
			} else if unicode.IsLower(r) {
				lowerCount++
			}
		}
	}
	return upperCount - lowerCount
}

func main() {
	// Test cases
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}