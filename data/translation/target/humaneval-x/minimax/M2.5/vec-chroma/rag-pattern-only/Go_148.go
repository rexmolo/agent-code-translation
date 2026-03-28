package main

func Bf(planet1, planet2 string) []string {
	planetNames := [8]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Check if either planet is not valid or if they are the same
	valid1 := false
	valid2 := false
	var planet1Index, planet2Index int

	for i, p := range planetNames {
		if p == planet1 {
			valid1 = true
			planet1Index = i
		}
		if p == planet2 {
			valid2 = true
			planet2Index = i
		}
	}

	if !valid1 || !valid2 || planet1Index == planet2Index {
		return []string{}
	}

	// Return planets between the two (exclusive)
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	}
	return planetNames[planet2Index+1 : planet1Index]
}