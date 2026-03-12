package main

import (
	"fmt"
	"unicode"
)

// calculateStrength computes the strength of a string based on letter casing.
func calculateStrength(s string) int {
	capCount := 0
	smCount := 0
	for _, r := range s {
		if unicode.IsUpper(r) {
			capCount++
		} else if unicode.IsLower(r) {
			smCount++
		}
	}
	return capCount - smCount
}

// StrongestExtension finds the extension with the highest strength and returns
// a formatted string "ClassName.StrongestExtensionName".
// The strength of an extension is calculated as the number of its uppercase
// letters minus the number of its lowercase letters.
// If there's a tie in strength, the extension that appears first in the list is chosen.
func StrongestExtension(className string, extensions []string) string {
	if len(extensions) == 0 {
		// The original Python code would panic on an empty list.
		// Assuming valid, non-empty input as per the examples.
		// A robust implementation might return an error or a default value.
		return className + "."
	}

	strongestExt := extensions[0]
	maxStrength := calculateStrength(extensions[0])

	// Start from the second element, as the first is already our candidate.
	for i := 1; i < len(extensions); i++ {
		ext := extensions[i]
		currentStrength := calculateStrength(ext)
		if currentStrength > maxStrength {
			maxStrength = currentStrength
			strongestExt = ext
		}
	}

	return fmt.Sprintf("%s.%s", className, strongestExt)
}
