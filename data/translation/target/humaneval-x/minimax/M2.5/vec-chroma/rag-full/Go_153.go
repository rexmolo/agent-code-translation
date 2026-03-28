package main

import (
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	// Initialize with the first extension
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

// calculateStrength computes the strength of an extension
// Strength = number of uppercase letters - number of lowercase letters
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
