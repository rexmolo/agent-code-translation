package main

import (
	"slices"
)

func Bf(planet1, planet2 string) []string {
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	
	idx1 := slices.Index(planets, planet1)
	idx2 := slices.Index(planets, planet2)
	
	// Return empty slice if either planet is not found or if they are the same
	if idx1 == -1 || idx2 == -1 || idx1 == idx2 {
		return []string{}
	}
	
	// Return planets between the two indices (exclusive), sorted by proximity to sun
	if idx1 < idx2 {
		return planets[idx1+1 : idx2]
	}
	return planets[idx2+1 : idx1]
}