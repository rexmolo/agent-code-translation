package main

import (
	"fmt"
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	// Start with the first extension as the strongest
	strong := extensions[0]
	myVal := calculateStrength(extensions[0])

	// Iterate through remaining extensions
	for _, s := range extensions[1:] {
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

	for _, c := range ext {
		if unicode.IsLetter(c) {
			if unicode.IsUpper(c) {
				upperCount++
			} else if unicode.IsLower(c) {
				lowerCount++
			}
		}
	}

	return upperCount - lowerCount
}

func main() {
	// Test examples
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}
