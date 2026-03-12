package main

import (
	"fmt"
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	strongest := extensions[0]

	// Calculate strength for the first extension (CAP - SM)
	myVal := calculateStrength(extensions[0])

	// Iterate through extensions to find the strongest
	for _, ext := range extensions {
		val := calculateStrength(ext)
		if val > myVal {
			strongest = ext
			myVal = val
		}
	}

	return className + "." + strongest
}

// calculateStrength returns CAP - SM (uppercase count minus lowercase count)
// Only alphabetic characters are considered
func calculateStrength(ext string) int {
	upper := 0
	lower := 0
	for _, r := range ext {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				upper++
			} else if unicode.IsLower(r) {
				lower++
			}
		}
	}
	return upper - lower
}

func main() {
	// Test cases
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}