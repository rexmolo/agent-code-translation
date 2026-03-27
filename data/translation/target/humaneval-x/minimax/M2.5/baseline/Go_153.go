package main

import "fmt"

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
	upper := 0
	lower := 0
	for _, ch := range ext {
		if ch >= 'A' && ch <= 'Z' {
			upper++
		} else if ch >= 'a' && ch <= 'z' {
			lower++
		}
	}
	return upper - lower
}

func main() {
	// Example tests
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
}