package main

import (
	"fmt"
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	// Handle empty extensions slice
	if len(extensions) == 0 {
		return className + "."
	}

	strong := extensions[0]
	myVal := calculateStrength(extensions[0])

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
	upper := 0
	lower := 0
	for _, ch := range ext {
		if unicode.IsUpper(ch) {
			upper++
		} else if unicode.IsLower(ch) {
			lower++
		}
	}
	return upper - lower
}

func main() {
	// Test examples
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}