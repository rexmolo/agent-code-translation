package main

import "fmt"

func Bf(planet1, planet2 string) []string {
	planetNames := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Check if either planet is invalid
	planet1Index := -1
	planet2Index := -1

	for i, p := range planetNames {
		if p == planet1 {
			planet1Index = i
		}
		if p == planet2 {
			planet2Index = i
		}
	}

	// Return empty slice if either planet is not valid or they are the same
	if planet1Index == -1 || planet2Index == -1 || planet1 == planet2 {
		return []string{}
	}

	// Return planets between the two (exclusive of both)
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}

func main() {
	// Test cases
	fmt.Println(Bf("Jupiter", "Neptune")) // [Saturn Uranus]
	fmt.Println(Bf("Earth", "Mercury"))   // [Venus]
	fmt.Println(Bf("Mercury", "Uranus")) // [Venus Earth Mars Jupiter Saturn]
	fmt.Println(Bf("Pluto", "Earth"))     // [] (invalid planet)
	fmt.Println(Bf("Earth", "Earth"))     // [] (same planet)
}