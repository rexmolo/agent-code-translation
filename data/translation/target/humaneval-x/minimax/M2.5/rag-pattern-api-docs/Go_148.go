package main

import "fmt"
import "slices"

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	planet1Index := slices.Index(planetNames, planet1)
	planet2Index := slices.Index(planetNames, planet2)

	// Return empty slice if planet names are invalid or the same
	if planet1Index == -1 || planet2Index == -1 || planet1 == planet2 {
		return []string{}
	}

	// Return planets between the two (exclusive), sorted by proximity to the sun
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}

func main() {
	// Test cases
	fmt.Println(Bf("Jupiter", "Neptune")) // [Saturn Uranus]
	fmt.Println(Bf("Earth", "Mercury"))   // [Venus]
	fmt.Println(Bf("Mercury", "Uranus"))  // [Venus Earth Mars Jupiter Saturn]
}