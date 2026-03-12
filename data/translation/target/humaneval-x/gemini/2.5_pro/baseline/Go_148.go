package main

var planetNames = []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
var planetIndices = make(map[string]int)

func init() {
	for i, name := range planetNames {
		planetIndices[name] = i
	}
}

func Bf(planet1, planet2 string) []string {
	index1, ok1 := planetIndices[planet1]
	index2, ok2 := planetIndices[planet2]

	if !ok1 || !ok2 || planet1 == planet2 {
		return nil
	}

	if index1 > index2 {
		index1, index2 = index2, index1
	}

	return planetNames[index1+1 : index2]
}