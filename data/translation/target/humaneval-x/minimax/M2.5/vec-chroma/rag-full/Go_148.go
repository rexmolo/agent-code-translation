package main

import "slices"

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Check if either planet is not valid, or if they are the same
	if slices.Index(planetNames, planet1) == -1 || slices.Index(planetNames, planet2) == -1 || planet1 == planet2 {
		return []string{}
	}

	planet1Index := slices.Index(planetNames, planet1)
	planet2Index := slices.Index(planetNames, planet2)

	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	} else {
		return planetNames[planet2Index+1 : planet1Index]
	}
}
