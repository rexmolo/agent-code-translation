package main

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Create a map for quick lookup
	validPlanets := make(map[string]bool)
	for _, p := range planetNames {
		validPlanets[p] = true
	}

	// Check if planets are valid or if they are the same
	if !validPlanets[planet1] || !validPlanets[planet2] || planet1 == planet2 {
		return []string{}
	}

	// Find indices of both planets
	var planet1Index, planet2Index int
	for i, p := range planetNames {
		if p == planet1 {
			planet1Index = i
		}
		if p == planet2 {
			planet2Index = i
		}
	}

	// Return planets between them (exclusive), sorted by proximity to the sun
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}