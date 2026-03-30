package main

import (
    "fmt"
    "slices"
)

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	planet1Index := slices.Index(planetNames, planet1)
	planet2Index := slices.Index(planetNames, planet2)

	// Return empty slice if either planet is not valid or if they are the same
	if planet1Index == -1 || planet2Index == -1 || planet1Index == planet2Index {
		return []string{}
	}

	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}

func main() {
	// Test cases from the examples
	fmt.Println(Bf("Jupiter", "Neptune")) // [Saturn Uranus]
	fmt.Println(Bf("Earth", "Mercury"))  // [Venus]
	fmt.Println(Bf("Mercury", "Uranus")) // [Venus Earth Mars Jupiter Saturn]
	
	// Edge cases
	fmt.Println(Bf("Invalid", "Earth")) // [] (invalid planet)
	fmt.Println(Bf("Earth", "Earth"))  // [] (same planet)
}
