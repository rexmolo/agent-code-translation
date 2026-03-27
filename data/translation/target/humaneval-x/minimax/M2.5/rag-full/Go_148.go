package main

import "fmt"

// Bf returns planets between two given planets in the solar system
func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	planet1Index := -1
	planet2Index := -1

	for i, name := range planetNames {
		if name == planet1 {
			planet1Index = i
		}
		if name == planet2 {
			planet2Index = i
		}
	}

	// Return empty slice if either planet is not valid or if they are the same
	if planet1Index == -1 || planet2Index == -1 || planet1 == planet2 {
		return []string{}
	}

	// Return planets between the two (exclusive)
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}

func main() {
	// Test examples
	fmt.Println(Bf("Jupiter", "Neptune")) // [Saturn Uranus]
	fmt.Println(Bf("Earth", "Mercury"))   // [Venus]
	fmt.Println(Bf("Mercury", "Uranus"))  // [Venus Earth Mars Jupiter Saturn]
	fmt.Println(Bf("Mars", "Mars"))       // []
	fmt.Println(Bf("Pluto", "Earth"))     // []
}