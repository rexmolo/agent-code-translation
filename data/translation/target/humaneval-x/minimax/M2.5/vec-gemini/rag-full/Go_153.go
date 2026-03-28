package main

import "fmt"

func StrongestExtension(class_name string, extensions []string) string {
	// Start with the first extension as the strongest
	strong := extensions[0]
	myVal := calculateStrength(extensions[0])

	// Iterate through all extensions to find the one with highest strength
	for _, s := range extensions {
		val := calculateStrength(s)
		if val > myVal {
			strong = s
			myVal = val
		}
	}

	return class_name + "." + strong
}

// calculateStrength computes CAP - SM for a given extension string
// where CAP is the number of uppercase letters and SM is the number of lowercase letters
func calculateStrength(s string) int {
	upperCount := 0
	lowerCount := 0

	for _, r := range s {
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
	result := StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"})
	fmt.Println(result)

	result2 := StrongestExtension("my_class", []string{"AA", "Be", "CC"})
	fmt.Println(result2)
}
