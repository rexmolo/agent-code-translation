package main

import "slices"

func LargestSmallestIntegers(lst []int) [2]interface{} {
	var negatives []int
	var positives []int

	for _, v := range lst {
		if v < 0 {
			negatives = append(negatives, v)
		} else if v > 0 {
			positives = append(positives, v)
		}
	}

	var largestNeg interface{} = nil
	if len(negatives) > 0 {
		largestNeg = slices.Max(negatives)
	}

	var smallestPos interface{} = nil
	if len(positives) > 0 {
		smallestPos = slices.Min(positives)
	}

	return [2]interface{}{largestNeg, smallestPos}
}