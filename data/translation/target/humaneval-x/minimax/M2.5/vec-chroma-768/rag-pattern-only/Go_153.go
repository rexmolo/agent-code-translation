package main

import (
	"fmt"
	"unicode"
)

func StrongestExtension(className string, extensions []string) string {
	strong := extensions[0]
	myVal := 0
	for _, c := range extensions[0] {
		if unicode.IsUpper(c) {
			myVal++
		} else if unicode.IsLower(c) {
			myVal--
		}
	}

	for _, s := range extensions {
		val := 0
		for _, c := range s {
			if unicode.IsUpper(c) {
				val++
			} else if unicode.IsLower(c) {
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
	fmt.Println(StrongestExtension("my_class", []string{"AA", "Be", "CC"}))
	fmt.Println(StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"}))
}
