package main

import "fmt"

func StrongestExtension(className string, extensions []string) string {
	strong := extensions[0]
	myVal := 0

	// Calculate strength for the first extension
	// Strength = number of uppercase letters - number of lowercase letters
	for _, c := range extensions[0] {
		if c >= 'A' && c <= 'Z' {
			myVal++
		} else if c >= 'a' && c <= 'z' {
			myVal--
		}
	}

	// Compare with other extensions
	for _, s := range extensions {
		val := 0
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				val++
			} else if c >= 'a' && c <= 'z' {
				val--
			}
		}
		if val > myVal {
			strong = s
			myVal = val
		}
	}

	return className + "." + strong
}

func main() {
	// Test examples
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
}