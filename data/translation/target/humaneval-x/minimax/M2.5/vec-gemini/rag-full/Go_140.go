package main

func FixSpaces(text string) string {
	newText := ""
	start := 0
	end := 0

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			end++
		} else {
			if end-start > 2 {
				newText += "-" + string(text[i])
			} else if end-start > 0 {
				for j := 0; j < end-start; j++ {
					newText += "_"
				}
				newText += string(text[i])
			} else {
				newText += string(text[i])
			}
			start = i + 1
			end = i + 1
		}
	}

	// Handle trailing spaces
	if end-start > 2 {
		newText += "-"
	} else if end-start > 0 {
		newText += "_"
	}

	return newText
}